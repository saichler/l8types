package nets

import (
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types/l8services"
	"github.com/saichler/l8types/go/types/l8sysconfig"
	"google.golang.org/protobuf/proto"

	"net"
)

func ExecuteProtocol(conn net.Conn, config *l8sysconfig.L8SysConfig, security ifs.ISecurityProvider) error {
	err := WriteEncrypted(conn, []byte(config.LocalUuid), config, security)
	if err != nil {
		conn.Close()
		return err
	}

	config.RemoteUuid, err = ReadEncrypted(conn, config, security)
	if err != nil {
		conn.Close()
		return err
	}

	forceExternal := "false"
	if config.ForceExternal {
		forceExternal = "true"
	}

	err = WriteEncrypted(conn, []byte(forceExternal), config, security)
	if err != nil {
		conn.Close()
		return err
	}

	forceExternal, err = ReadEncrypted(conn, config, security)
	if err != nil {
		conn.Close()
		return err
	}
	if forceExternal == "true" {
		config.ForceExternal = true
	}

	err = WriteEncrypted(conn, []byte(config.LocalAlias), config, security)
	if err != nil {
		conn.Close()
		return err
	}

	remoteAlias, err := ReadEncrypted(conn, config, security)
	if err != nil {
		conn.Close()
		return err
	}
	config.RemoteAlias = remoteAlias

	err = WriteEncrypted(conn, ServicesToBytes(config.Services), config, security)
	if err != nil {
		conn.Close()
		return err
	}

	services, err := ReadEncryptedBytes(conn, config, security)
	if err != nil {
		conn.Close()
		return err
	}
	config.Services = BytesToServices(services)

	return nil
}

func ServicesToBytes(services *l8services.L8Services) []byte {
	data, err := proto.Marshal(services)
	if err != nil {
		return []byte{}
	}
	return data
}

func BytesToServices(data []byte) *l8services.L8Services {
	services := &l8services.L8Services{}
	err := proto.Unmarshal(data, services)
	if err != nil {
		return nil
	}
	return services
}
