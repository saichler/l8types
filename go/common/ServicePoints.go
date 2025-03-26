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
	Handle(proto.Message, types.Action, IVirtualNetworkInterface, *types.Message, bool) Response
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(proto.Message, IVirtualNetworkInterface, *types.Message, bool) Response
	// Return the service point handler for the service name and area
	ServicePointHandler(string, int32) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(proto.Message, IResources) Response
	Put(proto.Message, IResources) Response
	Patch(proto.Message, IResources) Response
	Delete(proto.Message, IResources) Response
	GetCopy(proto.Message, IResources) Response
	Get(proto.Message, IResources) Response
	Failed(proto.Message, IResources, *types.Message) Response
	ServiceName() string
	ServiceModel() proto.Message
	EndPoint() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}
