package nets

import (
	"errors"
	"fmt"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types"
	"net"
)

// Write data to socket
func Write(data []byte, conn net.Conn, config *types.SysConfig) error {
	// If the connection is nil, return an error
	if conn == nil {
		return errors.New("no Connection Available")
	}
	// If the config is nil, error
	if config == nil {
		return errors.New("no Config Available")
	}
	if data == nil {
		return errors.New("no Data Available")
	}
	// Error is the data is too big
	if uint64(len(data)) > config.MaxDataSize {
		return fmt.Errorf("data size %d exceeds maximum %d", len(data), config.MaxDataSize)
	}
	// Write the size of the data
	_, e := conn.Write(ifs.Long2Bytes(int64(len(data))))
	if e != nil {
		return e
	}
	// Write the actual data
	_, e = conn.Write(data)
	return e
}

func WriteEncrypted(conn net.Conn, data []byte, config *types.SysConfig,
	securityProvider ifs.ISecurityProvider) error {
	encData, err := securityProvider.Encrypt(data)
	if err != nil {
		return err
	}
	return Write([]byte(encData), conn, config)
}
