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

package metadata

import "github.com/uber/cherami-thrift/.generated/go/metadata"
import "github.com/uber/cherami-thrift/.generated/go/shared"
import "github.com/stretchr/testify/mock"

// MetadataService is an autogenerated mock type for the MetadataService type
type MetadataService struct {
	mock.Mock
}

// CreateConsumerGroup provides a mock function with given fields: registerRequest
func (_m *MetadataService) CreateConsumerGroup(registerRequest *shared.CreateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(registerRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(*shared.CreateConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(registerRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.CreateConsumerGroupRequest) error); ok {
		r1 = rf(registerRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConsumerGroupExtent provides a mock function with given fields: request
func (_m *MetadataService) CreateConsumerGroupExtent(request *metadata.CreateConsumerGroupExtentRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.CreateConsumerGroupExtentRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDestination provides a mock function with given fields: createRequest
func (_m *MetadataService) CreateDestination(createRequest *shared.CreateDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(createRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(*shared.CreateDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(createRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.CreateDestinationRequest) error); ok {
		r1 = rf(createRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateExtent provides a mock function with given fields: request
func (_m *MetadataService) CreateExtent(request *shared.CreateExtentRequest) (*shared.CreateExtentResult_, error) {
	ret := _m.Called(request)

	var r0 *shared.CreateExtentResult_
	if rf, ok := ret.Get(0).(func(*shared.CreateExtentRequest) *shared.CreateExtentResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.CreateExtentResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.CreateExtentRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConsumerGroup provides a mock function with given fields: deleteRequest
func (_m *MetadataService) DeleteConsumerGroup(deleteRequest *shared.DeleteConsumerGroupRequest) error {
	ret := _m.Called(deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(*shared.DeleteConsumerGroupRequest) error); ok {
		r0 = rf(deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDestination provides a mock function with given fields: deleteRequest
func (_m *MetadataService) DeleteDestination(deleteRequest *shared.DeleteDestinationRequest) error {
	ret := _m.Called(deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(*shared.DeleteDestinationRequest) error); ok {
		r0 = rf(deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDestinationUUID provides a mock function with given fields: deleteRequest
func (_m *MetadataService) DeleteDestinationUUID(deleteRequest *metadata.DeleteDestinationUUIDRequest) error {
	ret := _m.Called(deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.DeleteDestinationUUIDRequest) error); ok {
		r0 = rf(deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HostAddrToUUID provides a mock function with given fields: hostAddr
func (_m *MetadataService) HostAddrToUUID(hostAddr string) (string, error) {
	ret := _m.Called(hostAddr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(hostAddr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConsumerGroups provides a mock function with given fields: listRequest
func (_m *MetadataService) ListConsumerGroups(listRequest *metadata.ListConsumerGroupRequest) (*metadata.ListConsumerGroupResult_, error) {
	ret := _m.Called(listRequest)

	var r0 *metadata.ListConsumerGroupResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListConsumerGroupRequest) *metadata.ListConsumerGroupResult_); ok {
		r0 = rf(listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListConsumerGroupResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListConsumerGroupRequest) error); ok {
		r1 = rf(listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDestinations provides a mock function with given fields: listRequest
func (_m *MetadataService) ListDestinations(listRequest *shared.ListDestinationsRequest) (*shared.ListDestinationsResult_, error) {
	ret := _m.Called(listRequest)

	var r0 *shared.ListDestinationsResult_
	if rf, ok := ret.Get(0).(func(*shared.ListDestinationsRequest) *shared.ListDestinationsResult_); ok {
		r0 = rf(listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListDestinationsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.ListDestinationsRequest) error); ok {
		r1 = rf(listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDestinationsByUUID provides a mock function with given fields: listRequest
func (_m *MetadataService) ListDestinationsByUUID(listRequest *shared.ListDestinationsByUUIDRequest) (*shared.ListDestinationsResult_, error) {
	ret := _m.Called(listRequest)

	var r0 *shared.ListDestinationsResult_
	if rf, ok := ret.Get(0).(func(*shared.ListDestinationsByUUIDRequest) *shared.ListDestinationsResult_); ok {
		r0 = rf(listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListDestinationsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.ListDestinationsByUUIDRequest) error); ok {
		r1 = rf(listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExtentsStats provides a mock function with given fields: request
func (_m *MetadataService) ListExtentsStats(request *shared.ListExtentsStatsRequest) (*shared.ListExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *shared.ListExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*shared.ListExtentsStatsRequest) *shared.ListExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.ListExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHosts provides a mock function with given fields: request
func (_m *MetadataService) ListHosts(request *metadata.ListHostsRequest) (*metadata.ListHostsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListHostsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListHostsRequest) *metadata.ListHostsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListHostsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListHostsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInputHostExtentsStats provides a mock function with given fields: request
func (_m *MetadataService) ListInputHostExtentsStats(request *metadata.ListInputHostExtentsStatsRequest) (*metadata.ListInputHostExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListInputHostExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListInputHostExtentsStatsRequest) *metadata.ListInputHostExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListInputHostExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListInputHostExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStoreExtentsStats provides a mock function with given fields: request
func (_m *MetadataService) ListStoreExtentsStats(request *metadata.ListStoreExtentsStatsRequest) (*metadata.ListStoreExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListStoreExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListStoreExtentsStatsRequest) *metadata.ListStoreExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListStoreExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListStoreExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveExtent provides a mock function with given fields: request
func (_m *MetadataService) MoveExtent(request *metadata.MoveExtentRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.MoveExtentRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadConsumerGroup provides a mock function with given fields: getRequest
func (_m *MetadataService) ReadConsumerGroup(getRequest *metadata.ReadConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(getRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(*metadata.ReadConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadConsumerGroupRequest) error); ok {
		r1 = rf(getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtent provides a mock function with given fields: request
func (_m *MetadataService) ReadConsumerGroupExtent(request *metadata.ReadConsumerGroupExtentRequest) (*metadata.ReadConsumerGroupExtentResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadConsumerGroupExtentResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadConsumerGroupExtentRequest) *metadata.ReadConsumerGroupExtentResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadConsumerGroupExtentRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtents provides a mock function with given fields: request
func (_m *MetadataService) ReadConsumerGroupExtents(request *metadata.ReadConsumerGroupExtentsRequest) (*metadata.ReadConsumerGroupExtentsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadConsumerGroupExtentsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadConsumerGroupExtentsRequest) *metadata.ReadConsumerGroupExtentsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadConsumerGroupExtentsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDestination provides a mock function with given fields: getRequest
func (_m *MetadataService) ReadDestination(getRequest *metadata.ReadDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(getRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(*metadata.ReadDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadDestinationRequest) error); ok {
		r1 = rf(getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadExtentStats provides a mock function with given fields: request
func (_m *MetadataService) ReadExtentStats(request *metadata.ReadExtentStatsRequest) (*metadata.ReadExtentStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadExtentStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadExtentStatsRequest) *metadata.ReadExtentStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadExtentStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadExtentStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadStoreExtentReplicaStats provides a mock function with given fields: request
func (_m *MetadataService) ReadStoreExtentReplicaStats(request *metadata.ReadStoreExtentReplicaStatsRequest) (*metadata.ReadStoreExtentReplicaStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadStoreExtentReplicaStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadStoreExtentReplicaStatsRequest) *metadata.ReadStoreExtentReplicaStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadStoreExtentReplicaStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadStoreExtentReplicaStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterHostUUID provides a mock function with given fields: request
func (_m *MetadataService) RegisterHostUUID(request *metadata.RegisterHostUUIDRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.RegisterHostUUIDRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SealExtent provides a mock function with given fields: request
func (_m *MetadataService) SealExtent(request *metadata.SealExtentRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.SealExtentRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAckOffset provides a mock function with given fields: request
func (_m *MetadataService) SetAckOffset(request *metadata.SetAckOffsetRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.SetAckOffsetRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOutputHost provides a mock function with given fields: request
func (_m *MetadataService) SetOutputHost(request *metadata.SetOutputHostRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.SetOutputHostRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UUIDToHostAddr provides a mock function with given fields: hostUUID
func (_m *MetadataService) UUIDToHostAddr(hostUUID string) (string, error) {
	ret := _m.Called(hostUUID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(hostUUID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsumerGroup provides a mock function with given fields: updateRequest
func (_m *MetadataService) UpdateConsumerGroup(updateRequest *shared.UpdateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(updateRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(*shared.UpdateConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.UpdateConsumerGroupRequest) error); ok {
		r1 = rf(updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsumerGroupExtentStatus provides a mock function with given fields: request
func (_m *MetadataService) UpdateConsumerGroupExtentStatus(request *metadata.UpdateConsumerGroupExtentStatusRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.UpdateConsumerGroupExtentStatusRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDestination provides a mock function with given fields: updateRequest
func (_m *MetadataService) UpdateDestination(updateRequest *shared.UpdateDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(updateRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(*shared.UpdateDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.UpdateDestinationRequest) error); ok {
		r1 = rf(updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDestinationDLQCursors provides a mock function with given fields: updateRequest
func (_m *MetadataService) UpdateDestinationDLQCursors(updateRequest *metadata.UpdateDestinationDLQCursorsRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(updateRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(*metadata.UpdateDestinationDLQCursorsRequest) *shared.DestinationDescription); ok {
		r0 = rf(updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.UpdateDestinationDLQCursorsRequest) error); ok {
		r1 = rf(updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExtentReplicaStats provides a mock function with given fields: request
func (_m *MetadataService) UpdateExtentReplicaStats(request *metadata.UpdateExtentReplicaStatsRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.UpdateExtentReplicaStatsRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExtentStats provides a mock function with given fields: request
func (_m *MetadataService) UpdateExtentStats(request *metadata.UpdateExtentStatsRequest) (*metadata.UpdateExtentStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.UpdateExtentStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.UpdateExtentStatsRequest) *metadata.UpdateExtentStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.UpdateExtentStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.UpdateExtentStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStoreExtentReplicaStats provides a mock function with given fields: request
func (_m *MetadataService) UpdateStoreExtentReplicaStats(request *metadata.UpdateStoreExtentReplicaStatsRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metadata.UpdateStoreExtentReplicaStatsRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
