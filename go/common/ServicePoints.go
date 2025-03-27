package common

import (
	"github.com/saichler/types/go/types"
)

// Add a bool for transaction
type IServicePoints interface {
	// Register A Service Point, handler + service area
	RegisterServicePoint(IServicePointHandler, int32) error
	// Handle a message and forward to the handler
	Handle(IMObjects, types.Action, IVirtualNetworkInterface, *types.Message, bool) IMObjects
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IMObjects, IVirtualNetworkInterface, *types.Message, bool) IMObjects
	// Return the service point handler for the service name and area
	ServicePointHandler(string, int32) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(IMObjects, IResources) IMObjects
	Put(IMObjects, IResources) IMObjects
	Patch(IMObjects, IResources) IMObjects
	Delete(IMObjects, IResources) IMObjects
	GetCopy(IMObjects, IResources) IMObjects
	Get(IMObjects, IResources) IMObjects
	Failed(IMObjects, IResources, *types.Message) IMObjects
	ServiceName() string
	ServiceModel() IMObjects
	EndPoint() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}
