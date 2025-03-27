package common

import (
	"github.com/saichler/types/go/types"
)

type IVirtualNetworkInterface interface {
	Start()
	Shutdown()
	Name() string
	SendMessage([]byte) error
	// Unicast a message without expecting response
	Unicast(string, string, int32, types.Action, interface{}) error
	// Unicast a message expecting response
	Request(string, string, int32, types.Action, interface{}) IResponse
	// Reply to a Request
	Reply(*types.Message, IResponse) error
	// Multicast a message to all service name listeners, without expecting a response
	Multicast(string, int32, types.Action, interface{}) error
	// Single a message to ONLY ONE service provider of the group,
	// not expecting a response. Provider is chosen by residency to the requester.
	Single(string, int32, types.Action, interface{}) error
	// SingleRequest same as single but expecting a response
	SingleRequest(string, int32, types.Action, interface{}) IResponse
	// Leader Same as SingleRequest but sending always to the leader.
	Leader(string, int32, types.Action, interface{}) IResponse
	Forward(*types.Message, string) IResponse
	API(string, int32) API
	Resources() IResources
}

type API interface {
	Post(interface{}) IResponse
	Put(interface{}) IResponse
	Patch(interface{}) IResponse
	Delete(interface{}) IResponse
	Get(string) IResponse
}

type IDatatListener interface {
	ShutdownVNic(IVirtualNetworkInterface)
	HandleData([]byte, IVirtualNetworkInterface)
	Failed([]byte, IVirtualNetworkInterface, string)
}

type IResponse interface {
	Add(interface{})
	SetError(error)
	Elems() []interface{}
	Elem() interface{}
	Error() error
	ToProto() *types.Response
}

func NewVNicConfig(maxDataSize uint64, txQueueSize, rxQueueSize uint64, vNetPort uint32) *types.VNicConfig {
	mc := &types.VNicConfig{
		MaxDataSize: maxDataSize,
		TxQueueSize: txQueueSize,
		RxQueueSize: rxQueueSize,
		VnetPort:    vNetPort,
	}
	return mc
}
