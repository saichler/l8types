package common

import (
	"github.com/saichler/types/go/types"
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

type IVirtualNetworkInterface interface {
	Start()
	Shutdown()
	Name() string
	SendMessage([]byte) error
	// Unicast a message without expecting response
	Unicast(string, string, uint16, Action, interface{}) error
	// Unicast a message expecting response
	Request(string, string, uint16, Action, interface{}) IElements
	// Reply to a Request
	Reply(IMessage, IElements) error
	// Multicast a message to all service name listeners, without expecting a response
	Multicast(string, uint16, Action, interface{}) error
	// Single a message to ONLY ONE service provider of the group,
	// not expecting a response. Provider is chosen by residency to the requester.
	Single(string, uint16, Action, interface{}) (string, error)
	// SingleRequest same as single but expecting a response
	SingleRequest(string, uint16, Action, interface{}) IElements
	// Leader Same as SingleRequest but sending always to the leader.
	Leader(string, uint16, Action, interface{}) IElements
	Forward(IMessage, string) IElements
	API(string, uint16) API
	Resources() IResources
	NotifyServiceAdded() error
	NotifyServiceRemoved(string, uint16) error
	PropertyChangeNotification(*types.NotificationSet)
	WaitForConnection()
}

type API interface {
	Post(interface{}) IElements
	Put(interface{}) IElements
	Patch(interface{}) IElements
	Delete(interface{}) IElements
	Get(string) IElements
}

type IDatatListener interface {
	ShutdownVNic(IVirtualNetworkInterface)
	HandleData([]byte, IVirtualNetworkInterface)
	Failed([]byte, IVirtualNetworkInterface, string)
}

func NewVNicConfig(maxDataSize uint64, txQueueSize, rxQueueSize uint64, vNetPort uint32) *types.SysConfig {
	mc := &types.SysConfig{
		MaxDataSize: maxDataSize,
		TxQueueSize: txQueueSize,
		RxQueueSize: rxQueueSize,
		VnetPort:    vNetPort,
	}
	return mc
}
