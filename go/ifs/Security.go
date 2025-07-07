package ifs

import (
	"errors"
	"github.com/saichler/l8types/go/types"
	"net"
	"plugin"
)

type ISecurityProvider interface {
	Authenticate(string, string) string
	Message(string) (*Message, error)

	CanDial(string, uint32) (net.Conn, error)
	CanAccept(net.Conn) error
	ValidateConnection(net.Conn, *types.SysConfig) error

	Encrypt([]byte) (string, error)
	Decrypt(string) ([]byte, error)

	CanDoAction(Action, IElements, string, string, ...string) error
	ScopeView(IElements, string, string, ...string) IElements
}

type ISecurityProviderLoader interface {
	LoadSecurityProvider(IResources) (ISecurityProvider, error)
}

func LoadSecurityProvider(resources IResources) (ISecurityProvider, error) {
	loaderFile, err := plugin.Open("./loader.so")
	if err != nil {
		return nil, errors.New("failed to load security provider error #1")
	}
	loader, err := loaderFile.Lookup("Loader")
	if err != nil {
		return nil, errors.New("failed to load security provider plugin #2")
	}
	if loader == nil {
		return nil, errors.New("failed to load security provider plugin #3")
	}
	loaderInterface := *loader.(*ISecurityProviderLoader)
	securityLoader := loaderInterface.(ISecurityProviderLoader).(ISecurityProviderLoader)
	return securityLoader.LoadSecurityProvider(resources)
}
