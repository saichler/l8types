package ifs

import (
	"bytes"
	"fmt"
)

func (this *Message) Unmarshal(data []byte, resources IResources) (interface{}, error) {
	// Basic bounds checking for header
	if len(data) < pPriority+1 {
		return nil, fmt.Errorf("insufficient data for message header: got %d bytes, need at least %d", len(data), pPriority+1)
	}

	this.source = stringFromBytes(data, pSource, pVnet)
	this.vnet = stringFromBytes(data, pVnet, pDestination)
	this.destination = optimizedToDestination(data)
	this.serviceName = optimizedToServiceName(data)
	this.serviceArea = data[pServiceArea]
	this.priority = Priority(data[pPriority])

	body, err := resources.Security().Decrypt(string(data[pPriority+1:]))
	if err != nil {
		return nil, err
	}

	// Basic bounds checking for body
	if len(body) < pFailMessageSize+1 {
		return nil, fmt.Errorf("insufficient body data: got %d bytes, need at least %d", len(body), pFailMessageSize+1)
	}

	this.action, this.tr_state = ByteToActionState(body[pAction])
	this.aaaId = stringFromBytes(body, pAaaId, pSequence)
	this.sequence = getUInt32(body, pSequence)
	this.timeout = getUInt16(body, pTimeout)
	this.request, this.reply = BoolOf(body[pRequestReply])

	failMessageSize := int(body[pFailMessageSize])
	pDataSize := pFailMessage + failMessageSize
	this.failMessage = stringFromBytes(body, pFailMessage, pDataSize)

	pData := pDataSize + sUint32
	dataSize := int(getUInt32(body, pDataSize))
	pTrId := pData + dataSize
	this.data = stringFromBytes(body, pData, pTrId)

	if this.tr_state != Empty {
		pTrErrMsgSize := pTrId + sUuid
		this.tr_id = stringFromBytes(body, pTrId, pTrErrMsgSize)
		trErrMsgSize := int(body[pTrErrMsgSize])
		pTrErrMsg := pTrErrMsgSize + sByte
		pTrStartTime := pTrErrMsg + trErrMsgSize
		this.tr_errMsg = stringFromBytes(body, pTrErrMsg, pTrStartTime)
		this.tr_startTime = getInt64(body, pTrStartTime)
	}

	return nil, nil
}

func HeaderOf(data []byte) (string, string, string, string, byte, Priority) {
	var serviceArea byte
	var priority Priority
	
	// Bounds checking for direct byte access
	if len(data) > pServiceArea {
		serviceArea = data[pServiceArea]
	}
	if len(data) > pPriority {
		priority = Priority(data[pPriority])
	}
	
	return stringFromBytes(data, pSource, pVnet),
		stringFromBytes(data, pVnet, pDestination),
		optimizedToDestination(data),
		optimizedToServiceName(data),
		serviceArea,
		priority
}

func ToDestination(data []byte) string {
	if data[pDestination] != 0 && data[pDestination+1] != 0 {
		return string(data[pDestination:pServiceName])
	}
	return ""
}

func optimizedToDestination(data []byte) string {
	if len(data) <= pDestination+1 || len(data) < pServiceName {
		return ""
	}
	if data[pDestination] != 0 && data[pDestination+1] != 0 {
		return stringFromBytes(data, pDestination, pServiceName)
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

func optimizedToServiceName(data []byte) string {
	return nullTerminatedString(data, pServiceName, sServiceName)
}
