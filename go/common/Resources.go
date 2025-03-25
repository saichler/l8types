package common

import (
	"github.com/google/uuid"
	"github.com/saichler/types/go/types"
)

type IResources interface {
	Registry() IRegistry
	ServicePoints() IServicePoints
	Security() ISecurityProvider
	DataListener() IDatatListener
	Serializer(SerializerMode) ISerializer
	Logger() ILogger
	Config() *types.VNicConfig
	Introspector() IIntrospector
	AddService(string, int32)
}

func AddService(config *types.VNicConfig, serviceName string, serviceArea int32) {
	if config == nil {
		return
	}
	if config.LocalUuid == "" {
		config.LocalUuid = NewUuid()
	}
	if config.Services == nil {
		config.Services = &types.Services{}
	}
	if config.Services.ServiceToAreas == nil {
		config.Services.ServiceToAreas = make(map[string]*types.ServiceAreas)
	}
	_, ok := config.Services.ServiceToAreas[serviceName]
	if !ok {
		config.Services.ServiceToAreas[serviceName] = &types.ServiceAreas{}
		config.Services.ServiceToAreas[serviceName].Areas = make(map[int32]*types.ServiceAreaInfo)
	}
	config.Services.ServiceToAreas[serviceName].Areas[serviceArea] = &types.ServiceAreaInfo{Score: 0}
}

func NewUuid() string {
	return uuid.New().String()
}
