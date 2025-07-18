package ifs

import "github.com/saichler/l8types/go/types"

// Add a bool for transaction
type IServices interface {
	// Add a service point type so compiling will pull the code for it
	RegisterServiceHandlerType(IServiceHandler)
	// Activate a service point
	Activate(string, string, byte, IResources, IServiceCacheListener, ...interface{}) (IServiceHandler, error)
	DeActivate(string, byte, IResources, IServiceCacheListener) error
	// Handle a message and forward to the handler
	Handle(IElements, Action, IVNic, *Message) IElements
	TransactionHandle(IElements, Action, IVNic, *Message) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVNic, *Message, bool) IElements
	// Return the service point handler for the service name and area
	ServiceHandler(string, byte) (IServiceHandler, bool)
	// Register a distributed cache
	RegisterDistributedCache(cache IDistributedCache)
}

type IServiceHandler interface {
	Activate(string, byte, IResources, IServiceCacheListener, ...interface{}) error
	DeActivate() error
	Post(IElements, IVNic) IElements
	Put(IElements, IVNic) IElements
	Patch(IElements, IVNic) IElements
	Delete(IElements, IVNic) IElements
	GetCopy(IElements, IVNic) IElements
	Get(IElements, IVNic) IElements
	Failed(IElements, IVNic, *Message) IElements
	TransactionMethod() ITransactionMethod
	WebService() IWebService
}

type IServiceCacheListener interface {
	PropertyChangeNotification(*types.NotificationSet)
}

type IDistributedCache interface {
	Put(string, interface{}, ...bool) (*types.NotificationSet, error)
	Update(string, interface{}, ...bool) (*types.NotificationSet, error)
	Delete(string, ...bool) (*types.NotificationSet, error)
	Get(k string) interface{}
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
	ServiceName() string
	ServiceArea() byte
	Sync()
}

type ITransactionMethod interface {
	Replication() bool
	ReplicationCount() int
	KeyOf(IElements, IResources) string
}
