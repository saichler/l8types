package nets

import (
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types"
	"google.golang.org/protobuf/proto"
	"net"
)

// closeOnError is a helper function to close connection and return error
func closeOnError(conn net.Conn, err error) error {
	if err != nil {
		conn.Close()
	}
	return err
}

func ExecuteProtocol(conn net.Conn, config *types.SysConfig, security ifs.ISecurityProvider) error {
	// Exchange UUIDs
	if err := WriteEncrypted(conn, []byte(config.LocalUuid), config, security); err != nil {
		return closeOnError(conn, err)
	}
	
	remoteUuid, err := ReadEncrypted(conn, config, security)
	if err != nil {
		return closeOnError(conn, err)
	}
	config.RemoteUuid = remoteUuid

	// Exchange ForceExternal flags
	forceExternalStr := "false"
	if config.ForceExternal {
		forceExternalStr = "true"
	}

	if err := WriteEncrypted(conn, []byte(forceExternalStr), config, security); err != nil {
		return closeOnError(conn, err)
	}

	remoteForceExternal, err := ReadEncrypted(conn, config, security)
	if err != nil {
		return closeOnError(conn, err)
	}
	config.ForceExternal = remoteForceExternal == "true"

	// Exchange aliases
	if err := WriteEncrypted(conn, []byte(config.LocalAlias), config, security); err != nil {
		return closeOnError(conn, err)
	}

	remoteAlias, err := ReadEncrypted(conn, config, security)
	if err != nil {
		return closeOnError(conn, err)
	}
	config.RemoteAlias = remoteAlias

	// Exchange services
	if err := WriteEncrypted(conn, ServicesToBytes(config.Services), config, security); err != nil {
		return closeOnError(conn, err)
	}

	servicesBytes, err := ReadEncryptedBytes(conn, config, security)
	if err != nil {
		return closeOnError(conn, err)
	}
	config.Services = BytesToServices(servicesBytes)

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
