package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestMessageClone(t *testing.T) {
	// Create a message with all fields populated
	msg := &ifs.Message{}
	msg.SetSource("original-source")
	msg.SetVnet("original-vnet")
	msg.SetDestination("original-dest")
	msg.SetServiceName("original-service")
	msg.SetServiceArea(byte(3))
	msg.SetPriority(ifs.P1)
	msg.SetMulticastMode(ifs.M_All)
	msg.SetAction(ifs.POST)
	msg.SetSequence(uint32(100))
	msg.SetTimeout(uint16(60))
	msg.SetRequestReply(true, false)
	msg.SetFailMessage("original fail")
	msg.SetData([]byte("original data"))
	msg.SetAAAId("original-aaa")
	msg.SetTr_State(ifs.Running)
	msg.SetTr_Id("original-tr-id")
	msg.SetTr_ErrMsg("original tr error")
	msg.SetTr_Timeout(int64(3000))
	msg.SetTr_Replica(byte(7))
	msg.SetTr_IsReplica(true)

	// Test Clone
	clone := msg.Clone()
	if clone.Source() != msg.Source() {
		t.Errorf("Clone source mismatch: expected %s, got %s", msg.Source(), clone.Source())
	}
	if clone.Vnet() != msg.Vnet() {
		t.Errorf("Clone vnet mismatch: expected %s, got %s", msg.Vnet(), clone.Vnet())
	}
	if clone.Destination() != msg.Destination() {
		t.Errorf("Clone destination mismatch: expected %s, got %s", msg.Destination(), clone.Destination())
	}
	if clone.ServiceName() != msg.ServiceName() {
		t.Errorf("Clone service name mismatch: expected %s, got %s", msg.ServiceName(), clone.ServiceName())
	}
	if clone.ServiceArea() != msg.ServiceArea() {
		t.Errorf("Clone service area mismatch: expected %d, got %d", msg.ServiceArea(), clone.ServiceArea())
	}
	if clone.Priority() != msg.Priority() {
		t.Errorf("Clone priority mismatch: expected %v, got %v", msg.Priority(), clone.Priority())
	}
	if clone.MulticastMode() != msg.MulticastMode() {
		t.Errorf("Clone multicast mode mismatch: expected %v, got %v", msg.MulticastMode(), clone.MulticastMode())
	}
	if clone.Action() != msg.Action() {
		t.Errorf("Clone action mismatch: expected %v, got %v", msg.Action(), clone.Action())
	}
	if clone.Sequence() != msg.Sequence() {
		t.Errorf("Clone sequence mismatch: expected %d, got %d", msg.Sequence(), clone.Sequence())
	}
	if clone.Timeout() != msg.Timeout() {
		t.Errorf("Clone timeout mismatch: expected %d, got %d", msg.Timeout(), clone.Timeout())
	}
	if clone.Request() != msg.Request() || clone.Reply() != msg.Reply() {
		t.Errorf("Clone request/reply mismatch")
	}
	if clone.FailMessage() != msg.FailMessage() {
		t.Errorf("Clone fail message mismatch: expected %s, got %s", msg.FailMessage(), clone.FailMessage())
	}
	if string(clone.Data()) != string(msg.Data()) {
		t.Errorf("Clone data mismatch: expected %s, got %s", string(msg.Data()), string(clone.Data()))
	}
	if clone.AAAId() != msg.AAAId() {
		t.Errorf("Clone AAA ID mismatch: expected %s, got %s", msg.AAAId(), clone.AAAId())
	}
	if clone.Tr_State() != msg.Tr_State() {
		t.Errorf("Clone tr state mismatch: expected %v, got %v", msg.Tr_State(), clone.Tr_State())
	}
	if clone.Tr_Id() != msg.Tr_Id() {
		t.Errorf("Clone tr id mismatch: expected %s, got %s", msg.Tr_Id(), clone.Tr_Id())
	}
	if clone.Tr_ErrMsg() != msg.Tr_ErrMsg() {
		t.Errorf("Clone tr error mismatch: expected %s, got %s", msg.Tr_ErrMsg(), clone.Tr_ErrMsg())
	}
	// Note: timing attributes are auto-set by SetTr_State, so they should be cloned
	if clone.Tr_Timeout() != msg.Tr_Timeout() {
		t.Errorf("Clone tr timeout mismatch: expected %d, got %d", msg.Tr_Timeout(), clone.Tr_Timeout())
	}
	if clone.Tr_Replica() != msg.Tr_Replica() {
		t.Errorf("Clone tr replica mismatch: expected %d, got %d", msg.Tr_Replica(), clone.Tr_Replica())
	}
	if clone.Tr_IsReplica() != msg.Tr_IsReplica() {
		t.Errorf("Clone tr isReplica mismatch: expected %v, got %v", msg.Tr_IsReplica(), clone.Tr_IsReplica())
	}
}

func TestMessageCloneReply(t *testing.T) {
	// Create a message
	msg := &ifs.Message{}
	msg.SetSource("original-source")
	msg.SetVnet("original-vnet")
	msg.SetDestination("original-dest")
	msg.SetServiceName("test-service")
	msg.SetServiceArea(byte(2))
	msg.SetPriority(ifs.P4)
	msg.SetMulticastMode(ifs.M_Unicast)
	msg.SetAction(ifs.GET)
	msg.SetSequence(uint32(200))
	msg.SetTimeout(uint16(45))
	msg.SetRequestReply(true, false)
	msg.SetData([]byte("test data"))
	msg.SetTr_State(ifs.Running)
	msg.SetTr_Replica(byte(2))
	msg.SetTr_IsReplica(false)

	// Test CloneReply
	localUuid := "local-uuid"
	remoteUuid := "remote-uuid"
	reply := msg.CloneReply(localUuid, remoteUuid)

	// Verify reply-specific changes
	if reply.Source() != localUuid {
		t.Errorf("CloneReply source should be local UUID: expected %s, got %s", localUuid, reply.Source())
	}
	if reply.Vnet() != remoteUuid {
		t.Errorf("CloneReply vnet should be remote UUID: expected %s, got %s", remoteUuid, reply.Vnet())
	}
	if reply.Destination() != msg.Source() {
		t.Errorf("CloneReply destination should be original source: expected %s, got %s", msg.Source(), reply.Destination())
	}
	if reply.Action() != ifs.Reply {
		t.Errorf("CloneReply action should be Reply: expected %v, got %v", ifs.Reply, reply.Action())
	}
	if !reply.Reply() {
		t.Error("CloneReply reply flag should be true")
	}
	if reply.Request() {
		t.Error("CloneReply request flag should be false")
	}

	// Verify other fields are preserved
	if reply.ServiceName() != msg.ServiceName() {
		t.Errorf("CloneReply service name mismatch: expected %s, got %s", msg.ServiceName(), reply.ServiceName())
	}
	if reply.Sequence() != msg.Sequence() {
		t.Errorf("CloneReply sequence mismatch: expected %d, got %d", msg.Sequence(), reply.Sequence())
	}
	if reply.Tr_Replica() != msg.Tr_Replica() {
		t.Errorf("CloneReply tr replica mismatch: expected %d, got %d", msg.Tr_Replica(), reply.Tr_Replica())
	}
	if reply.Tr_IsReplica() != msg.Tr_IsReplica() {
		t.Errorf("CloneReply tr isReplica mismatch: expected %v, got %v", msg.Tr_IsReplica(), reply.Tr_IsReplica())
	}
}

func TestMessageCloneFail(t *testing.T) {
	// Create a message
	msg := &ifs.Message{}
	msg.SetSource("original-source")
	msg.SetVnet("original-vnet")
	msg.SetDestination("original-dest")
	msg.SetServiceName("test-service")
	msg.SetAction(ifs.PUT)
	msg.SetSequence(uint32(300))
	msg.SetTr_State(ifs.Failed)
	msg.SetTr_Replica(byte(4))
	msg.SetTr_IsReplica(true)

	// Test CloneFail
	failMessage := "Operation failed"
	remoteUuid := "remote-uuid"
	failClone := msg.CloneFail(failMessage, remoteUuid)

	// Verify fail-specific changes
	if failClone.Source() != msg.Destination() {
		t.Errorf("CloneFail source should be original destination: expected %s, got %s", msg.Destination(), failClone.Source())
	}
	if failClone.Vnet() != remoteUuid {
		t.Errorf("CloneFail vnet should be remote UUID: expected %s, got %s", remoteUuid, failClone.Vnet())
	}
	if failClone.Destination() != msg.Source() {
		t.Errorf("CloneFail destination should be original source: expected %s, got %s", msg.Source(), failClone.Destination())
	}
	if failClone.FailMessage() != failMessage {
		t.Errorf("CloneFail fail message mismatch: expected %s, got %s", failMessage, failClone.FailMessage())
	}

	// Verify other fields are preserved
	if failClone.ServiceName() != msg.ServiceName() {
		t.Errorf("CloneFail service name mismatch: expected %s, got %s", msg.ServiceName(), failClone.ServiceName())
	}
	if failClone.Action() != msg.Action() {
		t.Errorf("CloneFail action mismatch: expected %v, got %v", msg.Action(), failClone.Action())
	}
	if failClone.Sequence() != msg.Sequence() {
		t.Errorf("CloneFail sequence mismatch: expected %d, got %d", msg.Sequence(), failClone.Sequence())
	}
	if failClone.Tr_Replica() != msg.Tr_Replica() {
		t.Errorf("CloneFail tr replica mismatch: expected %d, got %d", msg.Tr_Replica(), failClone.Tr_Replica())
	}
	if failClone.Tr_IsReplica() != msg.Tr_IsReplica() {
		t.Errorf("CloneFail tr isReplica mismatch: expected %v, got %v", msg.Tr_IsReplica(), failClone.Tr_IsReplica())
	}
}
