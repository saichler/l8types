package ifs

import (
	"github.com/saichler/l8types/go/types/l8api"
	"github.com/saichler/l8types/go/types/l8notify"
	"github.com/saichler/l8types/go/types/l8services"
)

// Add a bool for transaction
type IServices interface {
	// Add a service point type so compiling will pull the code for it
	RegisterServiceHandlerType(IServiceHandler)
	// Activate a service point
	Activate(*ServiceLevelAgreement, IVNic) (IServiceHandler, error)
	DeActivate(string, byte, IResources, IServiceCacheListener) error
	// Handle a message and forward to the handler
	Handle(IElements, Action, *Message, IVNic) IElements
	TransactionHandle(IElements, Action, *Message, IVNic) IElements
	// Handle a notification message, massage it to a change set and forward to the handler
	Notify(IElements, IVNic, *Message, bool) IElements
	// Return the service point handler for the service name and area
	ServiceHandler(string, byte) (IServiceHandler, bool)
	//The list of existing services
	Services() *l8services.L8Services
	GetLeader(string, byte) string
	GetParticipants(string, byte) map[string]byte
	RoundRobinParticipants(string, byte, int) map[string]byte
	TriggerElections(nic IVNic)
}

type IServiceHandler interface {
	Activate(*ServiceLevelAgreement, IVNic) error
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

type IServiceHandlerCache interface {
	IServiceHandler
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
	All() map[string]interface{}
	Fetch(int, int, IQuery) ([]interface{}, *l8api.L8Counts)
	AddCountFunc(string, func(interface{}) (bool, string))
	ServiceName() string
	ServiceArea() byte
	Size() int
}

type IMapReduceService interface {
	Merge(map[string]IElements) IElements
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
	Fetch(int, int, IQuery) ([]interface{}, *l8api.L8Counts)
	AddCountFunc(string, func(interface{}) (bool, string))
}

type IReplicationCache interface {
	Post(interface{}, int) error
	Put(interface{}, int) error
	Patch(interface{}, int) error
	Delete(interface{}, int) error
	Get(interface{}, int) (interface{}, error)
}

type ITransactionConfig interface {
	Replication() bool
	ReplicationCount() int
	KeyOf(IElements, IResources) string
	Voter() bool
}
