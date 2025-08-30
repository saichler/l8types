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

	this.action, this.tr_state = ByteToActionState(body[pAction])
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

	if this.tr_state != Empty {
		pTrErrMsgSize := pTrId + sUuid
		this.tr_id = unsafeString(body[pTrId:pTrErrMsgSize])
		trErrMsgSize := int(body[pTrErrMsgSize])
		pTrErrMsg := pTrErrMsgSize + sByte
		pTrStartTime := pTrErrMsg + trErrMsgSize
		this.tr_errMsg = unsafeString(body[pTrErrMsg:pTrStartTime])
		this.tr_startTime = Bytes2Long(body[pTrStartTime:])
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
