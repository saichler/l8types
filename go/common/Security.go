package common

import (
	"bytes"
	"errors"
	"github.com/saichler/types/go/types"
	"net"
	"os"
	"plugin"
)

type ISecurityProvider interface {
	CanDial(string, uint32) (net.Conn, error)
	CanAccept(net.Conn) error
	ValidateConnection(net.Conn, *types.SysConfig) error

	Encrypt([]byte) (string, error)
	Decrypt(string) ([]byte, error)

	CanDoAction(Action, IElements, string, string, ...string) error
	ScopeView(IElements, string, string, ...string) IElements
	Authenticate(string, string, ...string) string
}

func LoadSecurityProvider(soFileName string) (ISecurityProvider, error) {
	path := SeekResource("./../../", soFileName)
	if path == "" {
		panic("Could not find " + soFileName)
	}
	securityProviderPlugin, err := plugin.Open(path)
	if err != nil {
		return nil, errors.New("failed to load security provider plugin #1 " + path + " " + err.Error())
	}
	securityProvider, err := securityProviderPlugin.Lookup("SecurityProvider")
	if err != nil {
		return nil, errors.New("failed to load security provider plugin #2")
	}
	if securityProvider == nil {
		return nil, errors.New("failed to load security provider plugin #3")
	}
	providerInterface := *securityProvider.(*ISecurityProvider)
	return providerInterface.(ISecurityProvider), nil
}

func SeekResource(path string, filename string) string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return ""
	}
	if fileInfo.Name() == filename {
		return path
	}
	if fileInfo.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			return ""
		}
		for _, file := range files {
			found := SeekResource(pathOf(path, file), filename)
			if found != "" {
				return found
			}
		}
	}
	return ""
}

func pathOf(path string, file os.DirEntry) string {
	buff := bytes.Buffer{}
	buff.WriteString(path)
	buff.WriteString("/")
	buff.WriteString(file.Name())
	return buff.String()
}
