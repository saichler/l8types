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
	PPriority    = pServiceArea + sByte

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
	pTrErrMsgSize := pTrId + sUuid
	pTrErrMsg := pTrErrMsgSize + sByte
	pTrStartTime := pTrErrMsg + trErrMsgSize
	pEnd := pTrStartTime + 8

	var bodySize int
	if this.tr_state == Empty {
		bodySize = pTrId
	} else {
		bodySize = pEnd
	}

	totalSize := PPriority + sByte + bodySize
	result := make([]byte, totalSize)

	header := result[:PPriority+sByte]
	copy(header[pSource:pVnet], this.source)
	copy(header[pVnet:pDestination], this.vnet)
	copy(header[pDestination:pServiceName], this.destination)
	copy(header[pServiceName:pServiceArea], this.serviceName)
	header[pServiceArea] = this.serviceArea
	header[PPriority] = priorityMulticastModeToByte(this.priority, this.multicastMode)

	body := result[PPriority+sByte:]
	body[pAction] = actionStateToByte(this.action, this.tr_state)
	copy(body[pAaaId:pSequence], this.aaaId)
	copy(body[pSequence:pTimeout], UInt322Bytes(this.sequence))
	copy(body[pTimeout:pRequestReply], UInt162Bytes(this.timeout))
	body[pRequestReply] = Bools(this.request, this.reply)
	body[pFailMessageSize] = byte(failMessageSize)
	copy(body[pFailMessage:pDataSize], this.failMessage)
	copy(body[pDataSize:pData], UInt322Bytes(uint32(dataSize)))
	copy(body[pData:pTrId], this.data)

	if this.tr_state != Empty {
		copy(body[pTrId:pTrErrMsgSize], this.tr_id)
		body[pTrErrMsgSize] = byte(trErrMsgSize)
		copy(body[pTrErrMsg:pTrStartTime], this.tr_errMsg)
		copy(body[pTrStartTime:pEnd], Long2Bytes(this.tr_startTime))
	}

	bodyEnc, err := resources.Security().Encrypt(body)
	if err != nil {
		return nil, err
	}

	headerSize := PPriority + sByte
	finalData := make([]byte, headerSize+len(bodyEnc))
	copy(finalData[:headerSize], header)
	copy(finalData[headerSize:], bodyEnc)

	return finalData, nil
}
