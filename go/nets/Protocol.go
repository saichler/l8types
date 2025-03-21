package nets

import (
	"github.com/saichler/types/go/common"
	"github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
	"net"
)

func ExecuteProtocol(conn net.Conn, config *types.VNicConfig, security common.ISecurityProvider) error {
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

func ServicesToBytes(services *types.Services) []byte {
	data, err := proto.Marshal(services)
	if err != nil {
		return []byte{}
	}
	return data
}

func BytesToServices(data []byte) *types.Services {
	services := &types.Services{}
	err := proto.Unmarshal(data, services)
	if err != nil {
		return nil
	}
	return services
}
