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
	pTrState         = pAction + sByte
	pAaaId           = pTrState + sByte
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
	pTrCreated := pTrErrMsg + trErrMsgSize
	pTrQueued := pTrCreated + 8
	pTrRunning := pTrQueued + 8
	pTrEnd := pTrRunning + 8
	pTrTimeout := pTrEnd + 8
	pTrReplica := pTrTimeout + 8
	pEnd := pTrReplica + sByte

	var bodySize int
	if this.tr_state == NotATransaction {
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
	body[pAction] = byte(this.action)
	body[pTrState] = byte(this.tr_state)
	copy(body[pAaaId:pSequence], this.aaaId)
	copy(body[pSequence:pTimeout], UInt322Bytes(this.sequence))
	copy(body[pTimeout:pRequestReply], UInt162Bytes(this.timeout))
	body[pRequestReply] = Bools(this.request, this.reply)
	body[pFailMessageSize] = byte(failMessageSize)
	copy(body[pFailMessage:pDataSize], this.failMessage)
	copy(body[pDataSize:pData], UInt322Bytes(uint32(dataSize)))
	copy(body[pData:pTrId], this.data)

	if this.tr_state != NotATransaction {
		copy(body[pTrId:pTrErrMsgSize], this.tr_id)
		body[pTrErrMsgSize] = byte(trErrMsgSize)
		copy(body[pTrErrMsg:pTrCreated], this.tr_errMsg)
		copy(body[pTrCreated:pTrQueued], Long2Bytes(this.tr_created))
		copy(body[pTrQueued:pTrRunning], Long2Bytes(this.tr_queued))
		copy(body[pTrRunning:pTrEnd], Long2Bytes(this.tr_running))
		copy(body[pTrEnd:pTrTimeout], Long2Bytes(this.tr_end))
		copy(body[pTrTimeout:pTrReplica], Long2Bytes(this.tr_timeout))
		body[pTrReplica] = this.tr_replica
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
