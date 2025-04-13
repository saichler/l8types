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
	SysConfig() *types.SysConfig
	Introspector() IIntrospector
	AddService(string, int32)
}

func AddService(sysConfig *types.SysConfig, serviceName string, serviceArea int32) {
	if sysConfig == nil {
		return
	}
	if sysConfig.LocalUuid == "" {
		sysConfig.LocalUuid = NewUuid()
	}
	if sysConfig.Services == nil {
		sysConfig.Services = &types.Services{}
	}
	if sysConfig.Services.ServiceToAreas == nil {
		sysConfig.Services.ServiceToAreas = make(map[string]*types.ServiceAreas)
	}
	_, ok := sysConfig.Services.ServiceToAreas[serviceName]
	if !ok {
		sysConfig.Services.ServiceToAreas[serviceName] = &types.ServiceAreas{}
		sysConfig.Services.ServiceToAreas[serviceName].Areas = make(map[int32]*types.ServiceAreaInfo)
	}
	sysConfig.Services.ServiceToAreas[serviceName].Areas[serviceArea] = &types.ServiceAreaInfo{Score: 0}
}

func RemoveService(sysConfig *types.SysConfig, serviceName string, serviceArea int32) {
	if sysConfig == nil {
		return
	}
	if sysConfig.LocalUuid == "" {
		sysConfig.LocalUuid = NewUuid()
	}
	if sysConfig.Services == nil {
		sysConfig.Services = &types.Services{}
	}
	if sysConfig.Services.ServiceToAreas == nil {
		sysConfig.Services.ServiceToAreas = make(map[string]*types.ServiceAreas)
	}
	_, ok := sysConfig.Services.ServiceToAreas[serviceName]
	if !ok {
		return
	}
	delete(sysConfig.Services.ServiceToAreas[serviceName].Areas, serviceArea)
	if len(sysConfig.Services.ServiceToAreas[serviceName].Areas) == 0 {
		delete(sysConfig.Services.ServiceToAreas, serviceName)
	}
}

func NewUuid() string {
	return uuid.New().String()
}
