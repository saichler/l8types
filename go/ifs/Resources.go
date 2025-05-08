package ifs

import (
	"github.com/google/uuid"
	"github.com/saichler/l8types/go/types"
)

type IResources interface {
	Registry() IRegistry
	Services() IServices
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
		sysConfig.Services.ServiceToAreas[serviceName].Areas = make(map[int32]bool)
	}
	sysConfig.Services.ServiceToAreas[serviceName].Areas[serviceArea] = true
}

func RemoveService(services *types.Services, serviceName string, serviceArea int32) {
	if services == nil {
		return
	}
	if services.ServiceToAreas == nil {
		return
	}
	_, ok := services.ServiceToAreas[serviceName]
	if !ok {
		return
	}
	delete(services.ServiceToAreas[serviceName].Areas, serviceArea)
	if len(services.ServiceToAreas[serviceName].Areas) == 0 {
		delete(services.ServiceToAreas, serviceName)
	}
}

func NewUuid() string {
	return uuid.New().String()
}
