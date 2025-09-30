package ifs

import (
	"github.com/google/uuid"
	"github.com/saichler/l8types/go/types/l8api"
	"github.com/saichler/l8types/go/types/l8services"
	"github.com/saichler/l8types/go/types/l8sysconfig"
)

type IResources interface {
	Registry() IRegistry
	Services() IServices
	Security() ISecurityProvider
	DataListener() IDatatListener
	Serializer(SerializerMode) ISerializer
	Logger() ILogger
	SysConfig() *l8sysconfig.L8SysConfig
	Introspector() IIntrospector
	AddService(string, int32)
	Set(interface{})
	Copy(IResources)
	DefaultUser() *l8api.AuthUser
}

func AddService(sysConfig *l8sysconfig.L8SysConfig, serviceName string, serviceArea int32) {
	if sysConfig == nil {
		return
	}
	if sysConfig.LocalUuid == "" {
		sysConfig.LocalUuid = NewUuid()
	}
	if sysConfig.Services == nil {
		sysConfig.Services = &l8services.L8Services{}
	}
	if sysConfig.Services.ServiceToAreas == nil {
		sysConfig.Services.ServiceToAreas = make(map[string]*l8services.L8ServiceAreas)
	}
	_, ok := sysConfig.Services.ServiceToAreas[serviceName]
	if !ok {
		sysConfig.Services.ServiceToAreas[serviceName] = &l8services.L8ServiceAreas{}
		sysConfig.Services.ServiceToAreas[serviceName].Areas = make(map[int32]bool)
	}
	sysConfig.Services.ServiceToAreas[serviceName].Areas[serviceArea] = true
}

func RemoveService(services *l8services.L8Services, serviceName string, serviceArea int32) {
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
