package common

import (
	"github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
)

// Add a bool for transaction
type IServicePoints interface {
	// Register A Service Point, handler + service area
	RegisterServicePoint(IServicePointHandler, int32) error
	// Handle a message and forward to the handler
	Handle(proto.Message, types.Action, IVirtualNetworkInterface, *types.Message, bool) IResponse
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(proto.Message, IVirtualNetworkInterface, *types.Message, bool) IResponse
	// Return the service point handler for the service name and area
	ServicePointHandler(string, int32) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(proto.Message, IResources) IResponse
	Put(proto.Message, IResources) IResponse
	Patch(proto.Message, IResources) IResponse
	Delete(proto.Message, IResources) IResponse
	GetCopy(proto.Message, IResources) IResponse
	Get(proto.Message, IResources) IResponse
	Failed(proto.Message, IResources, *types.Message) IResponse
	ServiceName() string
	ServiceModel() proto.Message
	EndPoint() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}
