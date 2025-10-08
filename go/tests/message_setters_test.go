package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestMessageSetters(t *testing.T) {
	msg := &ifs.Message{}

	// Test SetSource
	msg.SetSource("test-source")
	if msg.Source() != "test-source" {
		t.Errorf("Expected source 'test-source', got '%s'", msg.Source())
	}

	// Test SetVnet
	msg.SetVnet("test-vnet")
	if msg.Vnet() != "test-vnet" {
		t.Errorf("Expected vnet 'test-vnet', got '%s'", msg.Vnet())
	}

	// Test SetDestination
	msg.SetDestination("test-dest")
	if msg.Destination() != "test-dest" {
		t.Errorf("Expected destination 'test-dest', got '%s'", msg.Destination())
	}

	// Test SetServiceName
	msg.SetServiceName("test-service")
	if msg.ServiceName() != "test-service" {
		t.Errorf("Expected service name 'test-service', got '%s'", msg.ServiceName())
	}

	// Test SetServiceArea
	msg.SetServiceArea(byte(5))
	if msg.ServiceArea() != byte(5) {
		t.Errorf("Expected service area 5, got %d", msg.ServiceArea())
	}

	// Test SetSequence
	msg.SetSequence(uint32(12345))
	if msg.Sequence() != uint32(12345) {
		t.Errorf("Expected sequence 12345, got %d", msg.Sequence())
	}

	// Test SetPriority
	msg.SetPriority(ifs.P1)
	if msg.Priority() != ifs.P1 {
		t.Errorf("Expected priority P1, got %v", msg.Priority())
	}

	// Test SetMulticastMode
	msg.SetMulticastMode(ifs.M_All)
	if msg.MulticastMode() != ifs.M_All {
		t.Errorf("Expected multicast mode M_All, got %v", msg.MulticastMode())
	}

	// Test SetAction
	msg.SetAction(ifs.POST)
	if msg.Action() != ifs.POST {
		t.Errorf("Expected action POST, got %v", msg.Action())
	}

	// Test SetTimeout
	msg.SetTimeout(uint16(30))
	if msg.Timeout() != uint16(30) {
		t.Errorf("Expected timeout 30, got %d", msg.Timeout())
	}

	// Test SetRequestReply
	msg.SetRequestReply(true, false)
	if !msg.Request() || msg.Reply() {
		t.Errorf("Expected request=true, reply=false, got request=%v, reply=%v", msg.Request(), msg.Reply())
	}

	msg.SetRequestReply(false, true)
	if msg.Request() || !msg.Reply() {
		t.Errorf("Expected request=false, reply=true, got request=%v, reply=%v", msg.Request(), msg.Reply())
	}

	// Test SetFailMessage
	msg.SetFailMessage("test error")
	if msg.FailMessage() != "test error" {
		t.Errorf("Expected fail message 'test error', got '%s'", msg.FailMessage())
	}

	// Test SetAAAId
	msg.SetAAAId("test-aaa-id")
	if msg.AAAId() != "test-aaa-id" {
		t.Errorf("Expected AAA ID 'test-aaa-id', got '%s'", msg.AAAId())
	}

	// Test SetData
	testData := []byte("test data")
	msg.SetData(testData)
	if string(msg.Data()) != "test data" {
		t.Errorf("Expected data 'test data', got '%s'", string(msg.Data()))
	}

	// Test SetTr_State
	msg.SetTr_State(ifs.Start)
	if msg.Tr_State() != ifs.Start {
		t.Errorf("Expected transaction state Start, got %v", msg.Tr_State())
	}

	// Test SetTr_Id
	msg.SetTr_Id("test-tr-id")
	if msg.Tr_Id() != "test-tr-id" {
		t.Errorf("Expected transaction ID 'test-tr-id', got '%s'", msg.Tr_Id())
	}

	// Test SetTr_ErrMsg
	msg.SetTr_ErrMsg("test tr error")
	if msg.Tr_ErrMsg() != "test tr error" {
		t.Errorf("Expected transaction error 'test tr error', got '%s'", msg.Tr_ErrMsg())
	}

	// Test SetTr_Created
	msg.SetTr_Created(int64(111111111))
	if msg.Tr_Created() != int64(111111111) {
		t.Errorf("Expected transaction created 111111111, got %d", msg.Tr_Created())
	}

	// Test SetTr_Queued
	msg.SetTr_Queued(int64(222222222))
	if msg.Tr_Queued() != int64(222222222) {
		t.Errorf("Expected transaction queued 222222222, got %d", msg.Tr_Queued())
	}

	// Test SetTr_Running
	msg.SetTr_Running(int64(333333333))
	if msg.Tr_Running() != int64(333333333) {
		t.Errorf("Expected transaction running 333333333, got %d", msg.Tr_Running())
	}

	// Test SetTr_End
	msg.SetTr_End(int64(444444444))
	if msg.Tr_End() != int64(444444444) {
		t.Errorf("Expected transaction end 444444444, got %d", msg.Tr_End())
	}

	// Test SetTr_Timeout
	msg.SetTr_Timeout(int64(987654321))
	if msg.Tr_Timeout() != int64(987654321) {
		t.Errorf("Expected transaction timeout 987654321, got %d", msg.Tr_Timeout())
	}
}
