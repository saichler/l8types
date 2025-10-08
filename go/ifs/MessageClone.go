package ifs

func (this *Message) Clone() *Message {
	clone := &Message{}
	clone.source = this.source
	clone.vnet = this.vnet
	clone.destination = this.destination
	clone.serviceName = this.serviceName
	clone.serviceArea = this.serviceArea
	clone.priority = this.priority
	clone.multicastMode = this.multicastMode
	clone.aaaId = this.aaaId
	clone.sequence = this.sequence
	clone.action = this.action
	clone.timeout = this.timeout
	clone.reply = this.reply
	clone.request = this.request
	clone.failMessage = this.failMessage
	clone.data = this.data
	clone.tr_id = this.tr_id
	clone.tr_state = this.tr_state
	clone.tr_errMsg = this.tr_errMsg
	clone.tr_created = this.tr_created
	clone.tr_queued = this.tr_queued
	clone.tr_running = this.tr_running
	clone.tr_end = this.tr_end
	clone.tr_timeout = this.tr_timeout
	return clone
}

func (this *Message) CloneReply(localUuid, remoteUuid string) *Message {
	clone := &Message{}
	clone.source = localUuid
	clone.vnet = remoteUuid
	clone.destination = this.source
	clone.serviceName = this.serviceName
	clone.serviceArea = this.serviceArea
	clone.priority = this.priority
	clone.multicastMode = this.multicastMode
	clone.aaaId = this.aaaId
	clone.sequence = this.sequence
	clone.action = Reply
	clone.timeout = this.timeout
	clone.reply = true
	clone.request = false
	clone.failMessage = this.failMessage
	clone.data = this.data
	clone.tr_id = this.tr_id
	clone.tr_state = this.tr_state
	clone.tr_errMsg = this.tr_errMsg
	clone.tr_created = this.tr_created
	clone.tr_queued = this.tr_queued
	clone.tr_running = this.tr_running
	clone.tr_end = this.tr_end
	clone.tr_timeout = this.tr_timeout
	return clone
}

func (this *Message) CloneFail(failMessage, remoteUuid string) *Message {
	clone := &Message{}
	clone.source = this.destination
	clone.vnet = remoteUuid
	clone.destination = this.source
	clone.serviceName = this.serviceName
	clone.serviceArea = this.serviceArea
	clone.priority = this.priority
	clone.multicastMode = this.multicastMode
	clone.aaaId = this.aaaId
	clone.sequence = this.sequence
	clone.action = this.action
	clone.timeout = this.timeout
	clone.reply = this.reply
	clone.request = this.request
	clone.failMessage = failMessage
	clone.data = this.data
	clone.tr_id = this.tr_id
	clone.tr_state = this.tr_state
	clone.tr_errMsg = this.tr_errMsg
	clone.tr_created = this.tr_created
	clone.tr_queued = this.tr_queued
	clone.tr_running = this.tr_running
	clone.tr_end = this.tr_end
	clone.tr_timeout = this.tr_timeout
	return clone
}
