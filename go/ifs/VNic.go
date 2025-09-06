package ifs

import (
	"github.com/saichler/l8types/go/types"
)

type NetworkMode int

const (
	NETWORK_NATIVE NetworkMode = 1
	NETWORK_DOCKER NetworkMode = 2
	NETWORK_K8s    NetworkMode = 3
)

const (
	SysMsg  = "sysMsg"
	SysArea = byte(99)
)

var networkMode NetworkMode = NETWORK_NATIVE

func SetNetworkMode(mode NetworkMode) {
	networkMode = mode
}

func NetworkMode_Native() bool {
	return networkMode == NETWORK_NATIVE
}

func NetworkMode_DOCKER() bool {
	return networkMode == NETWORK_DOCKER
}

func NetworkMode_K8s() bool {
	return networkMode == NETWORK_K8s
}

type IVNic interface {
	Start()
	Shutdown()
	Name() string
	SendMessage([]byte) error
	// Unicast a message without expecting response
	Unicast(string, string, byte, Action, interface{}) error
	// Request unicast a message expecting response
	Request(string, string, byte, Action, interface{}, ...string) IElements
	// Reply to a Request
	Reply(*Message, IElements) error
	// Multicast a message to all service name listeners, without expecting a response
	Multicast(string, byte, Action, interface{}) error
	// RoundRobin a message to ONLY ONE service provider of the group, in a round-robin fashion
	RoundRobin(string, byte, Action, interface{}) error
	// RoundRobinRequest a request to ONLY ONE service provider of the group, in a round-robin fashion
	RoundRobinRequest(string, byte, Action, interface{}) IElements
	// Proximity a message to ONLY ONE service provider of the group with a proximity of the provider to the sender
	Proximity(string, byte, Action, interface{}) error
	// Proximity a request to ONLY ONE service provider of the group with a proximity of the provider to the sender
	ProximityRequest(string, byte, Action, interface{}) IElements
	// Leader a message to ONLY ONE service provider leader of the group.
	Leader(string, byte, Action, interface{}) error
	// LeaderRequest a request to ONLY ONE service provider leader of the group.
	LeaderRequest(string, byte, Action, interface{}) IElements
	// Local a message to ONLY ONE service provider that resides in the same vnic.
	Local(string, byte, Action, interface{}) error
	// LocalRequest a request to ONLY ONE service provider that resides in the same vnic.
	LocalRequest(string, byte, Action, interface{}) IElements
	Forward(*Message, string) IElements
	ServiceAPI(string, byte) ServiceAPI
	Resources() IResources
	NotifyServiceAdded([]string, byte) error
	NotifyServiceRemoved(string, byte) error
	PropertyChangeNotification(*types.NotificationSet)
	WaitForConnection()
	Running() bool
	RegisterServiceBatch(string, byte, MulticastMode, int)
}

type ServiceAPI interface {
	Post(interface{}) IElements
	Put(interface{}) IElements
	Patch(interface{}) IElements
	Delete(interface{}) IElements
	Get(string) IElements
}

type IDatatListener interface {
	ShutdownVNic(nic IVNic)
	HandleData([]byte, IVNic)
	Failed([]byte, IVNic, string)
}
