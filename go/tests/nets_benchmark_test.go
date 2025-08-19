package tests

import (
	"bytes"
	"testing"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/nets"
	"github.com/saichler/l8types/go/types"
)

// Benchmark Write function
func BenchmarkWrite(b *testing.B) {
	config := &types.SysConfig{
		MaxDataSize: 1024 * 1024,
	}
	data := make([]byte, 1024) // 1KB test data
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		err := nets.Write(data, conn, config)
		if err != nil {
			b.Fatalf("Write failed: %v", err)
		}
	}
}

// Benchmark Read function
func BenchmarkRead(b *testing.B) {
	config := &types.SysConfig{
		MaxDataSize: 1024 * 1024,
	}
	testData := make([]byte, 1024) // 1KB test data
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Pre-create read buffer
	sizeBytes := ifs.Long2Bytes(int64(len(testData)))
	readData := append(sizeBytes, testData...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		conn.SetReadData(readData)
		
		result, err := nets.Read(conn, config)
		if err != nil {
			b.Fatalf("Read failed: %v", err)
		}
		if len(result) != len(testData) {
			b.Fatalf("Data length mismatch: expected %d, got %d", len(testData), len(result))
		}
	}
}

// Benchmark ReadSize function
func BenchmarkReadSize(b *testing.B) {
	config := &types.SysConfig{
		MaxDataSize: 1024 * 1024,
	}
	testData := make([]byte, 1024) // 1KB test data
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		conn.SetReadData(testData)
		
		result, err := nets.ReadSize(len(testData), conn, config)
		if err != nil {
			b.Fatalf("ReadSize failed: %v", err)
		}
		if len(result) != len(testData) {
			b.Fatalf("Data length mismatch: expected %d, got %d", len(testData), len(result))
		}
	}
}

// Benchmark WriteEncrypted function
func BenchmarkWriteEncrypted(b *testing.B) {
	config := &types.SysConfig{
		MaxDataSize: 1024 * 1024,
	}
	security := &MockSecurityProviderNets{}
	data := make([]byte, 1024) // 1KB test data
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		err := nets.WriteEncrypted(conn, data, config, security)
		if err != nil {
			b.Fatalf("WriteEncrypted failed: %v", err)
		}
	}
}

// Benchmark ReadEncrypted function
func BenchmarkReadEncrypted(b *testing.B) {
	config := &types.SysConfig{
		MaxDataSize: 1024 * 1024,
	}
	security := &MockSecurityProviderNets{}
	testData := make([]byte, 1024) // 1KB test data
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Pre-create encrypted read buffer
	sizeBytes := ifs.Long2Bytes(int64(len(testData)))
	readData := append(sizeBytes, testData...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		conn.SetReadData(readData)
		
		result, err := nets.ReadEncrypted(conn, config, security)
		if err != nil {
			b.Fatalf("ReadEncrypted failed: %v", err)
		}
		if len(result) != len(testData) {
			b.Fatalf("Data length mismatch: expected %d, got %d", len(testData), len(result))
		}
	}
}

// Benchmark ServicesToBytes function
func BenchmarkServicesToBytes(b *testing.B) {
	services := &types.Services{
		ServiceToAreas: map[string]*types.ServiceAreas{
			"service1": {
				Areas: map[int32]bool{1: true, 2: false, 3: true},
			},
			"service2": {
				Areas: map[int32]bool{4: true, 5: true},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := nets.ServicesToBytes(services)
		if len(data) == 0 {
			b.Fatal("ServicesToBytes returned empty data")
		}
	}
}

// Benchmark BytesToServices function
func BenchmarkBytesToServices(b *testing.B) {
	services := &types.Services{
		ServiceToAreas: map[string]*types.ServiceAreas{
			"service1": {
				Areas: map[int32]bool{1: true, 2: false, 3: true},
			},
			"service2": {
				Areas: map[int32]bool{4: true, 5: true},
			},
		},
	}
	data := nets.ServicesToBytes(services)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := nets.BytesToServices(data)
		if result == nil {
			b.Fatal("BytesToServices returned nil")
		}
	}
}

// Benchmark ExecuteProtocol function
func BenchmarkExecuteProtocol(b *testing.B) {
	security := &MockSecurityProviderNets{}
	config := &types.SysConfig{
		LocalUuid:     "local-uuid-123456789012345678901234567890",
		LocalAlias:    "local-alias",
		ForceExternal: false,
		MaxDataSize:   1024 * 1024,
		Services: &types.Services{
			ServiceToAreas: map[string]*types.ServiceAreas{
				"test-service": {
					Areas: map[int32]bool{1: true},
				},
			},
		},
	}

	// Pre-create protocol responses
	remoteUuid := "remote-uuid-456789012345678901234567890"
	remoteAlias := "remote-alias"
	forceExternal := "false"

	var responses [][]byte

	// Remote UUID response
	uuidSize := ifs.Long2Bytes(int64(len(remoteUuid)))
	responses = append(responses, append(uuidSize, []byte(remoteUuid)...))

	// ForceExternal response
	feSize := ifs.Long2Bytes(int64(len(forceExternal)))
	responses = append(responses, append(feSize, []byte(forceExternal)...))

	// Remote alias response
	aliasSize := ifs.Long2Bytes(int64(len(remoteAlias)))
	responses = append(responses, append(aliasSize, []byte(remoteAlias)...))

	// Services response
	servicesData := nets.ServicesToBytes(config.Services)
	servicesSize := ifs.Long2Bytes(int64(len(servicesData)))
	responses = append(responses, append(servicesSize, servicesData...))

	allResponses := bytes.Join(responses, []byte{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := NewMockConn()
		conn.SetReadData(allResponses)

		// Reset config for each iteration
		config.RemoteUuid = ""
		config.RemoteAlias = ""

		err := nets.ExecuteProtocol(conn, config, security)
		if err != nil {
			b.Fatalf("ExecuteProtocol failed: %v", err)
		}
	}
}