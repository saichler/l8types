package tests

import (
	"testing"

	"github.com/saichler/l8types/go/nets"
	"github.com/saichler/l8types/go/types/l8sysconfig"
)

func TestWriteNilConnection(t *testing.T) {
	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 1024000,
	}
	data := []byte("test data")

	err := nets.Write(data, nil, config)
	if err == nil {
		t.Error("Expected error when writing to nil connection")
	}
	if err.Error() != "no Connection Available" {
		t.Errorf("Expected 'no Connection Available' error, got: %s", err.Error())
	}
}

func TestWriteNilConfig(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	data := []byte("test data")
	err := nets.Write(data, conn, nil)
	if err == nil {
		t.Error("Expected error when writing with nil config")
	}
	if err.Error() != "no Config Available" {
		t.Errorf("Expected 'no Config Available' error, got: %s", err.Error())
	}
}

func TestWriteNilData(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 1024000,
	}

	err := nets.Write(nil, conn, config)
	if err == nil {
		t.Error("Expected error when writing nil data")
	}
	if err.Error() != "no Data Available" {
		t.Errorf("Expected 'no Data Available' error, got: %s", err.Error())
	}
}

func TestWriteDataTooLarge(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 100,
	}
	data := make([]byte, 200)

	err := nets.Write(data, conn, config)
	if err == nil {
		t.Error("Expected error when data exceeds max size")
	}
	if err.Error() != "data is larger than MAX size allowed" {
		t.Errorf("Expected 'data is larger than MAX size allowed' error, got: %s", err.Error())
	}
}

func TestReadNilConnection(t *testing.T) {
	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 1024000,
	}

	_, err := nets.Read(nil, config)
	if err == nil {
		t.Error("Expected error when reading from nil connection")
	}
	if err.Error() != "no Connection Available" {
		t.Errorf("Expected 'no Connection Available' error, got: %s", err.Error())
	}
}

func TestReadNilConfig(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	_, err := nets.Read(conn, nil)
	if err == nil {
		t.Error("Expected error when reading with nil config")
	}
	if err.Error() != "no Config Available" {
		t.Errorf("Expected 'no Config Available' error, got: %s", err.Error())
	}
}

func TestWriteEncryptedError(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 1024000,
	}
	security := &MockSecurityProviderNets{encryptError: true}
	data := []byte("test data")

	err := nets.WriteEncrypted(conn, data, config, security)
	if err == nil {
		t.Error("Expected error when encryption fails")
	}
}

func TestWriteEncryptedSuccess(t *testing.T) {
	conn := NewMockConn()
	defer conn.Close()

	config := &l8sysconfig.L8SysConfig{
		MaxDataSize: 1024000,
	}
	security := &MockSecurityProviderNets{}
	data := []byte("test data")

	err := nets.WriteEncrypted(conn, data, config, security)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}
