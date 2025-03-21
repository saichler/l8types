package common

import (
	"github.com/saichler/types/go/types"
)

/* Cast mode
enum CastMode {
Invalid_Cast_mode = 0;
All = 1;
Single = 2;
Leader = 3;
}*/

type IVirtualNetworkInterface interface {
	Start()
	Shutdown()
	Name() string
	SendMessage([]byte) error
	// Unicast a message without expecting response
	Unicast(types.Action, string, string, int32, interface{}) error
	// Unicast a message expecting response
	Request(types.Action, string, string, int32, interface{}) (interface{}, error)
	// Reply to a Request
	Reply(*types.Message, interface{}) error
	// Multicast a message to all service name listeners, without expecting a response
	Multicast(types.Action, string, int32, interface{}) error
	// Single a message to ONLY ONE service provider of the group,
	// not expecting a response. Provider is chosen by residency to the requester.
	Single(types.Action, string, int32, interface{}) error
	// SingleRequest same as single but expecting a response
	SingleRequest(types.Action, string, int32, interface{}) (interface{}, error)
	// Leader Same as SingleRequest but sending always to the leader.
	Leader(types.Action, string, int32, interface{}) (interface{}, error)
	Forward(*types.Message, string) (interface{}, error)
	API(int32) API
	Resources() IResources
}

type API interface {
	Post(interface{}) (interface{}, error)
	Put(interface{}) (interface{}, error)
	Patch(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
	Get(string) (interface{}, error)
}

type IDatatListener interface {
	ShutdownVNic(IVirtualNetworkInterface)
	HandleData([]byte, IVirtualNetworkInterface)
	Failed([]byte, IVirtualNetworkInterface, string)
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
