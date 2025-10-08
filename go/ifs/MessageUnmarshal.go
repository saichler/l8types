package ifs

import (
	"unsafe"
)

func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (this *Message) Unmarshal(data []byte, resources IResources) (interface{}, error) {

	this.source = unsafeString(data[pSource:pVnet])
	this.vnet = unsafeString(data[pVnet:pDestination])
	this.destination = ToDestination(data)
	this.serviceName = ToServiceName(data)
	this.serviceArea = data[pServiceArea]
	this.priority, this.multicastMode = ByteToPriorityMulticastMode(data[PPriority])

	body, err := resources.Security().Decrypt(string(data[PPriority+1:]))
	if err != nil {
		return nil, err
	}

	this.action = Action(body[pAction])
	this.tr_state = TransactionState(body[pTrState])
	this.aaaId = unsafeString(body[pAaaId:pSequence])
	this.sequence = Bytes2UInt32(body[pSequence:pTimeout])
	this.timeout = Bytes2UInt16(body[pTimeout:pRequestReply])
	this.request, this.reply = BoolOf(body[pRequestReply])

	failMessageSize := int(body[pFailMessageSize])
	pDataSize := pFailMessage + failMessageSize
	this.failMessage = unsafeString(body[pFailMessage:pDataSize])

	pData := pDataSize + sUint32
	dataSize := int(Bytes2UInt32(body[pDataSize:pData]))
	pTrId := pData + dataSize
	this.data = body[pData:pTrId]

	if this.tr_state != NotATransaction {
		pTrErrMsgSize := pTrId + sUuid
		this.tr_id = unsafeString(body[pTrId:pTrErrMsgSize])
		trErrMsgSize := int(body[pTrErrMsgSize])
		pTrErrMsg := pTrErrMsgSize + sByte
		pTrCreated := pTrErrMsg + trErrMsgSize
		pTrQueued := pTrCreated + 8
		pTrRunning := pTrQueued + 8
		pTrEnd := pTrRunning + 8
		pTrTimeout := pTrEnd + 8
		this.tr_errMsg = unsafeString(body[pTrErrMsg:pTrCreated])
		this.tr_created = Bytes2Long(body[pTrCreated:pTrQueued])
		this.tr_queued = Bytes2Long(body[pTrQueued:pTrRunning])
		this.tr_running = Bytes2Long(body[pTrRunning:pTrEnd])
		this.tr_end = Bytes2Long(body[pTrEnd:pTrTimeout])
		this.tr_timeout = Bytes2Long(body[pTrTimeout:])
	}

	return nil, nil
}

func HeaderOf(data []byte) (string, string, string, string, byte, Priority, MulticastMode) {
	return unsafeString(data[pSource:pVnet]),
		unsafeString(data[pVnet:pDestination]),
		ToDestination(data),
		ToServiceName(data),
		data[pServiceArea],
		Priority(data[PPriority] >> 4),
		MulticastMode(data[PPriority] & 0x0F)
}

func ToDestination(data []byte) string {
	if data[pDestination] != 0 && data[pDestination+1] != 0 {
		return unsafeString(data[pDestination:pServiceName])
	}
	return ""
}

func ToServiceName(data []byte) string {
	start := pServiceName
	end := start + sServiceName
	if end > len(data) {
		end = len(data)
	}

	for i := start; i < end; i++ {
		if data[i] == 0 {
			return unsafeString(data[start:i])
		}
	}
	return unsafeString(data[start:end])
}
