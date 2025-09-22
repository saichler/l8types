package ifs

import (
	"github.com/saichler/l8types/go/types/l8notify"
	"github.com/saichler/l8types/go/types/l8services"
)

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
	//The list of existing services
	Services() *l8services.L8Services
}

type IServiceHandler interface {
	Activate(string, byte, IResources, IServiceCacheListener, ...interface{}) error
	DeActivate() error
	Post(IElements, IVNic) IElements
	Put(IElements, IVNic) IElements
	Patch(IElements, IVNic) IElements
	Delete(IElements, IVNic) IElements
	Get(IElements, IVNic) IElements
	Failed(IElements, IVNic, *Message) IElements
	TransactionConfig() ITransactionConfig
	WebService() IWebService
}

type IServiceCacheListener interface {
	PropertyChangeNotification(set *l8notify.L8NotificationSet)
}

type IDistributedCache interface {
	Post(interface{}, ...bool) (*l8notify.L8NotificationSet, error)
	Put(interface{}, ...bool) (*l8notify.L8NotificationSet, error)
	Patch(interface{}, ...bool) (*l8notify.L8NotificationSet, error)
	Delete(interface{}, ...bool) (*l8notify.L8NotificationSet, error)
	Get(interface{}) (interface{}, error)
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
	ServiceName() string
	ServiceArea() byte
	Sync()
	Fetch(int, int, IQuery) []interface{}
	Stats() map[string]int32
	AddStatFunc(string, func(interface{}) bool)
}

type ITransactionConfig interface {
	Replication() bool
	ReplicationCount() int
	KeyOf(IElements, IResources) string
	ConcurrentGets() bool
}
