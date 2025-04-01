package common

import (
	"github.com/saichler/types/go/types"
)

// Add a bool for transaction
type IServicePoints interface {
	// Register A Service Point, handler + service area
	RegisterServicePoint(IServicePointHandler, int32) error
	// Handle a message and forward to the handler
	Handle(IElements, types.Action, IVirtualNetworkInterface, *types.Message, bool) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVirtualNetworkInterface, *types.Message, bool) IElements
	// Return the service point handler for the service name and area
	ServicePointHandler(string, int32) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Post(IElements, IResources) IElements
	Put(IElements, IResources) IElements
	Patch(IElements, IResources) IElements
	Delete(IElements, IResources) IElements
	GetCopy(IElements, IResources) IElements
	Get(string, IResources) IElements
	Failed(IElements, IResources, *types.Message) IElements
	ServiceName() string
	ServiceModel() IElements
	EndPoint() string
	Transactional() bool
	ReplicationCount() int
	ReplicationScore() int
}
