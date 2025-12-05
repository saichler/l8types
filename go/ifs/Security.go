package ifs

import (
	"errors"
	"fmt"
	"net"
	"plugin"

	"github.com/saichler/l8types/go/types/l8sysconfig"
)

type ISecurityProvider interface {
	Authenticate(string, string) (string, bool, bool, error)
	ValidateToken(string) (string, bool)

	Message(string) (*Message, error)

	CanDial(string, uint32) (net.Conn, error)
	CanAccept(net.Conn) error
	ValidateConnection(net.Conn, *l8sysconfig.L8SysConfig) error

	Encrypt([]byte) (string, error)
	Decrypt(string) ([]byte, error)

	CanDoAction(Action, IElements, string, string, ...string) error
	ScopeView(IElements, string, string, ...string) IElements

	TFASetup(string, IVNic) (string, []byte, error)
	TFAVerify(string, string, string, IVNic) error

	Captcha() []byte
	Register(string, string, string, IVNic) bool
}

type ISecurityProviderLoader interface {
	LoadSecurityProvider(...interface{}) (ISecurityProvider, error)
}

type ISecurityProviderActivate interface {
	Activate(IVNic)
}

func LoadSecurityProvider(args ...interface{}) (ISecurityProvider, error) {
	loaderFile, err := plugin.Open("/var/loader.so")
	if err != nil {
		fmt.Println("Failed to load security provider #1: ", err.Error())
		return nil, errors.New("Failed to load security provider #1: " + err.Error())
	}
	loader, err := loaderFile.Lookup("Loader")
	if err != nil {
		fmt.Println("Failed to load security provider #2: ", err.Error())
		return nil, errors.New("Failed to load security provider #2: " + err.Error())
	}
	if loader == nil {
		fmt.Println("Failed to load security provider #3: Nil Loader")
		return nil, errors.New("Failed to load security provider #3: Nil Loader")
	}
	loaderInterface := *loader.(*ISecurityProviderLoader)
	securityLoader := loaderInterface.(ISecurityProviderLoader).(ISecurityProviderLoader)
	return securityLoader.LoadSecurityProvider(args...)
}
