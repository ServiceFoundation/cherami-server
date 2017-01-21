// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package replicator

import (
	"sync"

	"github.com/uber-common/bark"
	"github.com/uber/cherami-server/common"
	"github.com/uber/cherami-server/common/metrics"
	storeStream "github.com/uber/cherami-server/stream"
	"github.com/uber/cherami-thrift/.generated/go/cherami"
	"github.com/uber/cherami-thrift/.generated/go/store"
)

type (
	outConnection struct {
		extUUID string
		stream  storeStream.BStoreOpenReadStreamOutCall
		msgsCh  chan *store.ReadMessageContent

		logger     bark.Logger
		m3Client   metrics.Client
		metricsTag int

		readMsgCountChannel chan int32    // channel to pass read msg count from readMsgStream to writeCreditsStream in order to issue more credits
		closeChannel        chan struct{} // channel to indicate the connection should be closed

		lk     sync.Mutex
		opened bool
		closed bool
	}
)

const (
	msgBufferSize = 10000

	initialCreditSize = 10000

	creditBatchSize = initialCreditSize / 10
)

func newOutConnection(extUUID string, stream storeStream.BStoreOpenReadStreamOutCall, logger bark.Logger, m3Client metrics.Client, metricsTag int) *outConnection {
	conn := &outConnection{
		extUUID:             extUUID,
		stream:              stream,
		msgsCh:              make(chan *store.ReadMessageContent, msgBufferSize),
		logger:              logger.WithField(common.TagExt, extUUID).WithField(`scope`, "outConnection"),
		m3Client:            m3Client,
		metricsTag:          metricsTag,
		readMsgCountChannel: make(chan int32, 10),
		closeChannel:        make(chan struct{}),
	}

	return conn
}

func (conn *outConnection) open() {
	conn.lk.Lock()
	defer conn.lk.Unlock()

	if !conn.opened {
		go conn.writeCreditsStream()
		go conn.readMsgStream()

		conn.opened = true
	}
	conn.logger.Info("out connection opened")
}

func (conn *outConnection) close() {
	conn.lk.Lock()
	defer conn.lk.Unlock()

	if !conn.closed {
		close(conn.closeChannel)
		conn.closed = true
	}

	conn.logger.Info("out connection closed")
}

func (conn *outConnection) writeCreditsStream() {
	defer conn.stream.Done()

	if err := conn.sendCredits(initialCreditSize); err != nil {
		conn.logger.Error(`error writing initial credits`)

		go conn.close()
		return
	}

	var numMsgsRead int32

	for {
		if numMsgsRead > 0 {
			if err := conn.sendCredits(numMsgsRead); err != nil {
				conn.logger.Error(`error sending credits`)

				go conn.close()
				return
			}
			numMsgsRead = 0
		} else {
			select {
			// Note: this will block until readMsgStream sends msg count to the channel, or the connection is closed
			case msgsRead := <-conn.readMsgCountChannel:
				numMsgsRead += msgsRead
			case <-conn.closeChannel:
				return
			}
		}
	}
}

func (conn *outConnection) readMsgStream() {
	// lastSeqNum is used to track whether our sequence numbers are
	// monotonically increasing
	// We initialize this to -1 to skip the first message check
	var lastSeqNum int64 = -1

	var sealMsgRead bool

	var numMsgsRead int32
	for {
		select {
		case <-conn.closeChannel:
			return
		default:
			if numMsgsRead >= creditBatchSize {
				select {
				case conn.readMsgCountChannel <- numMsgsRead:
					numMsgsRead = 0
				default:
					// Not the end of world if the channel is blocked
					conn.logger.WithField(`credit`, numMsgsRead).Info("readMsgStream: blocked sending credits; accumulating credits to send later")
				}
			}

			rmc, err := conn.stream.Read()
			if err != nil {
				conn.logger.WithField(common.TagErr, err).Error(`Error reading msg`)
				go conn.close()
				return
			}

			switch rmc.GetType() {
			case store.ReadMessageContentType_MESSAGE:
				msg := rmc.GetMessage()

				if sealMsgRead {
					conn.logger.WithFields(bark.Fields{
						"seqNum": msg.Message.GetSequenceNumber(),
					}).Error("regular message read after seal message")
					go conn.close()
					return
				}

				// Sequence number check to make sure we get monotonically increasing sequence number.
				if lastSeqNum+1 != msg.Message.GetSequenceNumber() && lastSeqNum != -1 {
					expectedSeqNum := 1 + lastSeqNum

					conn.logger.WithFields(bark.Fields{
						"seqNum":         msg.Message.GetSequenceNumber(),
						"expectedSeqNum": expectedSeqNum,
					}).Error("sequence number out of order")
					go conn.close()
					return
				}

				// update the lastSeqNum to this value
				lastSeqNum = msg.Message.GetSequenceNumber()

				conn.m3Client.IncCounter(conn.metricsTag, metrics.ReplicatorOutConnMsgRead)

				// now push msg to the msg channel (which will in turn be pushed to client)
				// Note this is a blocking call here
				select {
				case conn.msgsCh <- rmc:
					numMsgsRead++
				case <-conn.closeChannel:
					conn.logger.Info(`writing msg to the channel failed because of shutdown`)
					return
				}

			case store.ReadMessageContentType_SEALED:
				seal := rmc.GetSealed()
				conn.logger.WithField(`SequenceNumber`, seal.GetSequenceNumber()).Info(`extent sealed`)
				sealMsgRead = true

				// now push msg to the msg channel (which will in turn be pushed to client)
				// Note this is a blocking call here
				select {
				case conn.msgsCh <- rmc:
					numMsgsRead++
				case <-conn.closeChannel:
					conn.logger.Info(`writing msg to the channel failed because of shutdown`)
					return
				}

				return
			case store.ReadMessageContentType_ERROR:
				msgErr := rmc.GetError()
				conn.logger.WithField(`Message`, msgErr.GetMessage()).Error(`received error from reading msg`)
				go conn.close()
				return
			default:
				conn.logger.WithField(`Type`, rmc.GetType()).Error(`received ReadMessageContent with unrecognized type`)
			}
		}
	}
}

func (conn *outConnection) sendCredits(credits int32) error {
	cFlow := cherami.NewControlFlow()
	cFlow.Credits = common.Int32Ptr(credits)
	err := conn.stream.Write(cFlow)
	if err == nil {
		err = conn.stream.Flush()
	}

	conn.m3Client.AddCounter(conn.metricsTag, metrics.ReplicatorOutConnCreditsSent, int64(credits))

	return err
}
