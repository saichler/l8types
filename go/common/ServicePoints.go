package common

import "github.com/saichler/types/go/types"

// Add a bool for transaction
type IServicePoints interface {
	// Add a service point type so compiling will pull the code for it
	AddServicePointType(IServicePointHandler)
	// Activate a service point
	Activate(string, string, uint16, IResources, IServicePointCacheListener, ...interface{}) (IServicePointHandler, error)
	DeActivate(string, uint16, IResources, IServicePointCacheListener) error
	// Handle a message and forward to the handler
	Handle(IElements, Action, IVirtualNetworkInterface, IMessage, bool) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVirtualNetworkInterface, IMessage, bool) IElements
	// Return the service point handler for the service name and area
	ServicePointHandler(string, uint16) (IServicePointHandler, bool)
}

type IServicePointHandler interface {
	Activate(string, uint16, IResources, IServicePointCacheListener, ...interface{}) error
	DeActivate() error
	Post(IElements, IResources) IElements
	Put(IElements, IResources) IElements
	Patch(IElements, IResources) IElements
	Delete(IElements, IResources) IElements
	GetCopy(IElements, IResources) IElements
	Get(IElements, IResources) IElements
	Failed(IElements, IResources, IMessage) IElements
	TransactionMethod() ITransactionMethod
}

type IServicePointCacheListener interface {
	PropertyChangeNotification(*types.NotificationSet)
}

type IDistributedCache interface {
	Put(string, interface{}, ...bool) (*types.NotificationSet, error)
	Update(string, interface{}, ...bool) (*types.NotificationSet, error)
	Delete(string, ...bool) (*types.NotificationSet, error)
	Get(k string) interface{}
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
}

type ITransactionMethod interface {
	Replication() bool
	ReplicationCount() int
	KeyOf(IElements) string
}
