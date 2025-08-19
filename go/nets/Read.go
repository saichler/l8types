package nets

import (
	"errors"
	"fmt"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types"
	"io"
	"net"
)

// Read data from socket
func Read(conn net.Conn, config *types.SysConfig) ([]byte, error) {
	// If the connection is nil, return an error
	if conn == nil {
		return nil, errors.New("no Connection Available")
	}
	// If the config is nil, error
	if config == nil {
		return nil, errors.New("no Config Available")
	}
	// read 8 bytes, e.g. long, hinting of the size of the byte array
	sizebytes, err := ReadSize(8, conn, config)
	if sizebytes == nil || err != nil {
		return nil, err
	}
	// Translate the 8 byte array into int64
	size := ifs.Bytes2Long(sizebytes)
	// If the size is larger than the MAX Data Size, return an error
	// this is to protect against overflowing the buffers
	// When data to send is > the max data size, one needs to split the data into chunks at a higher level
	if uint64(size) > config.MaxDataSize {
		return nil, fmt.Errorf("data size %d exceeds maximum %d", size, config.MaxDataSize)
	}
	// Read the bunch of bytes according to the size from the socket
	data, err := ReadSize(int(size), conn, config)
	return data, err
}

func ReadSize(size int, conn net.Conn, config *types.SysConfig) ([]byte, error) {
	data := make([]byte, size)
	_, err := io.ReadFull(conn, data)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return data, nil
}

func ReadEncryptedBytes(conn net.Conn, config *types.SysConfig,
	securityProvider ifs.ISecurityProvider) ([]byte, error) {
	inData, err := Read(conn, config)
	if err != nil {
		conn.Close()
		return nil, err
	}

	decData, err := securityProvider.Decrypt(string(inData))
	if err != nil {
		conn.Close()
		return nil, err
	}
	return decData, nil
}

func ReadEncrypted(conn net.Conn, config *types.SysConfig,
	securityProvider ifs.ISecurityProvider) (string, error) {
	data, err := ReadEncryptedBytes(conn, config, securityProvider)
	return string(data), err
}
