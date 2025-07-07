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
	// Unicast a message expecting response
	Request(string, string, byte, Action, interface{}, ...string) IElements
	// Reply to a Request
	Reply(*Message, IElements) error
	// Multicast a message to all service name listeners, without expecting a response
	Multicast(string, byte, Action, interface{}) error
	// Single a message to ONLY ONE service provider of the group,
	// not expecting a response. Provider is chosen by residency to the requester.
	Single(string, byte, Action, interface{}) (string, error)
	// SingleRequest same as single but expecting a response
	SingleRequest(string, byte, Action, interface{}) IElements
	// Leader Same as SingleRequest but sending always to the leader.
	Leader(string, byte, Action, interface{}) IElements
	Forward(*Message, string) IElements
	ServiceAPI(string, byte) ServiceAPI
	Resources() IResources
	NotifyServiceAdded([]string, byte) error
	NotifyServiceRemoved(string, byte) error
	PropertyChangeNotification(*types.NotificationSet)
	WaitForConnection()
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
