package ifs

import "bytes"

func (this *Message) Unmarshal(data []byte, resources IResources) (interface{}, error) {

	this.source = string(data[pSource:pVnet])
	this.vnet = string(data[pVnet:pDestination])
	this.destination = ToDestination(data)
	this.serviceName = ToServiceName(data)
	this.serviceArea = data[pServiceArea]
	this.priority = Priority(data[pPriority])

	body, err := resources.Security().Decrypt(string(data[pPriority+1:]))
	if err != nil {
		return nil, err
	}

	this.action, this.tr_state = ByteToActionState(body[pAction])
	this.aaaId = string(body[pAaaId:pSequence])
	this.sequence = Bytes2UInt32(body[pSequence:pTimeout])
	this.timeout = Bytes2UInt16(body[pTimeout:pRequestReply])
	this.request, this.reply = BoolOf(body[pRequestReply])

	failMessageSize := int(body[pFailMessageSize])
	pDataSize := pFailMessage + failMessageSize
	this.failMessage = string(body[pFailMessage:pDataSize])

	pData := pDataSize + sUint32
	dataSize := int(Bytes2UInt32(body[pDataSize:pData]))
	pTrId := pData + dataSize
	this.data = string(body[pData:pTrId])

	if this.tr_state != Empty {
		pTrErrMsgSize := pTrId + sUuid
		this.tr_id = string(body[pTrId:pTrErrMsgSize])
		trErrMsgSize := int(body[pTrErrMsgSize])
		pTrErrMsg := pTrErrMsgSize + sByte
		pTrStartTime := pTrErrMsg + trErrMsgSize
		this.tr_errMsg = string(body[pTrErrMsg:pTrStartTime])
		this.tr_startTime = Bytes2Long(body[pTrStartTime:])
	}

	return nil, nil
}

func HeaderOf(data []byte) (string, string, string, string, byte, Priority) {
	return string(data[pSource:pVnet]),
		string(data[pVnet:pDestination]),
		ToDestination(data),
		ToServiceName(data),
		data[pServiceArea],
		Priority(data[pPriority])
}

func ToDestination(data []byte) string {
	if data[pDestination] != 0 && data[pDestination+1] != 0 {
		return string(data[pDestination:pServiceName])
	}
	return ""
}

func ToServiceName(data []byte) string {
	buff := bytes.Buffer{}
	for i := pServiceName; i < pServiceName+len(data); i++ {
		if data[i] != 0 {
			buff.WriteByte(data[i])
		} else {
			return buff.String()
		}
	}
	return buff.String()
}
