package ifs

const (
	sUuid        = 36
	sServiceName = 10
	sUint32      = 4
	sUint16      = 2
	sByte        = 1

	pSource      = 0
	pVnet        = pSource + sUuid
	pDestination = pVnet + sUuid
	pServiceName = pDestination + sUuid
	pServiceArea = pServiceName + sServiceName
	pPriority    = pServiceArea + sByte

	pAction          = 0
	pAaaId           = pAction + sByte
	pSequence        = pAaaId + sUuid
	pTimeout         = pSequence + sUint32
	pRequestReply    = pTimeout + sUint16
	pFailMessageSize = pRequestReply + sByte
	pFailMessage     = pFailMessageSize + sByte
)

func (this *Message) Marshal(any interface{}, resources IResources) ([]byte, error) {
	failMessageSize := len(this.failMessage)
	dataSize := len(this.data)
	trErrMsgSize := len(this.tr_errMsg)
	
	pDataSize := pFailMessage + failMessageSize
	pData := pDataSize + sUint32
	pTrId := pData + dataSize
	
	var bodySize int
	if this.tr_state == Empty {
		bodySize = pTrId
	} else {
		bodySize = pTrId + sUuid + sByte + trErrMsgSize + 8
	}
	
	body := make([]byte, bodySize)
	
	body[pAction] = actionStateToByte(this.action, this.tr_state)
	copy(body[pAaaId:pSequence], this.aaaId)
	putUInt32(body, pSequence, this.sequence)
	putUInt16(body, pTimeout, this.timeout)
	body[pRequestReply] = encodeBools(this.request, this.reply)
	body[pFailMessageSize] = byte(failMessageSize)
	copy(body[pFailMessage:pDataSize], this.failMessage)
	putUInt32(body, pDataSize, uint32(dataSize))
	copy(body[pData:pTrId], this.data)
	
	if this.tr_state != Empty {
		pTrErrMsgSize := pTrId + sUuid
		pTrErrMsg := pTrErrMsgSize + sByte
		pTrStartTime := pTrErrMsg + trErrMsgSize
		
		copy(body[pTrId:pTrErrMsgSize], this.tr_id)
		body[pTrErrMsgSize] = byte(trErrMsgSize)
		copy(body[pTrErrMsg:pTrStartTime], this.tr_errMsg)
		putInt64(body, pTrStartTime, this.tr_startTime)
	}
	
	bodyEnc, err := resources.Security().Encrypt(body)
	if err != nil {
		return nil, err
	}
	
	headerSize := pPriority + sByte
	result := make([]byte, headerSize+len(bodyEnc))
	
	copy(result[pSource:pVnet], this.source)
	copy(result[pVnet:pDestination], this.vnet)
	copy(result[pDestination:pServiceName], this.destination)
	copy(result[pServiceName:pServiceArea], this.serviceName)
	result[pServiceArea] = this.serviceArea
	result[pPriority] = byte(this.priority)
	
	copy(result[headerSize:], bodyEnc)
	
	return result, nil
}
