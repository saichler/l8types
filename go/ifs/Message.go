/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ifs

import "time"

type Message struct {
	source        string
	vnet          string
	destination   string
	serviceName   string
	serviceArea   byte
	priority      Priority
	multicastMode MulticastMode

	action      Action
	tr_state    TransactionState
	aaaId       string
	sequence    uint32
	timeout     uint16
	request     bool
	reply       bool
	failMessage string
	data        []byte

	tr_id        string
	tr_errMsg    string
	tr_created   int64
	tr_queued    int64
	tr_running   int64
	tr_end       int64
	tr_timeout   int64
	tr_replica   byte
	tr_isReplica bool
}

func (this *Message) Init(destination, serviceName string, serviceArea byte,
	priority Priority, multicastMode MulticastMode, action Action, source, vnet string, data []byte,
	isRequest, isReply bool, msgNum uint32,
	tr_state TransactionState, tr_id, tr_errMsg string, tr_created, tr_queued, tr_running, tr_end, tr_timeout int64, tr_replica byte, tr_isreplica bool) {
	this.destination = destination
	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.priority = priority
	this.multicastMode = multicastMode
	this.action = action
	this.source = source
	this.vnet = vnet
	this.data = data
	this.request = isRequest
	this.reply = isReply
	this.sequence = msgNum
	this.tr_state = tr_state
	this.tr_id = tr_id
	this.tr_errMsg = tr_errMsg
	this.tr_created = tr_created
	this.tr_queued = tr_queued
	this.tr_running = tr_running
	this.tr_end = tr_end
	this.tr_timeout = tr_timeout
	this.tr_replica = tr_replica
	this.tr_isReplica = tr_isreplica
}

//Getters

func (this *Message) Source() string {
	return this.source
}

func (this *Message) Vnet() string {
	return this.vnet
}

func (this *Message) Destination() string {
	return this.destination
}

func (this *Message) ServiceName() string {
	return this.serviceName
}

func (this *Message) ServiceArea() byte {
	return this.serviceArea
}

func (this *Message) Sequence() uint32 {
	return this.sequence
}

func (this *Message) Priority() Priority {
	return this.priority
}

func (this *Message) MulticastMode() MulticastMode {
	return this.multicastMode
}

func (this *Message) Action() Action {
	return this.action
}

func (this *Message) Timeout() uint16 {
	return this.timeout
}

func (this *Message) Request() bool {
	return this.request
}

func (this *Message) Reply() bool {
	return this.reply
}

func (this *Message) FailMessage() string {
	return this.failMessage
}

func (this *Message) Data() []byte {
	return this.data
}

func (this *Message) AAAId() string {
	return this.aaaId
}

func (this *Message) Tr_State() TransactionState {
	return this.tr_state
}

func (this *Message) Tr_Id() string {
	return this.tr_id
}

func (this *Message) Tr_ErrMsg() string {
	return this.tr_errMsg
}

func (this *Message) Tr_Created() int64 {
	return this.tr_created
}

func (this *Message) Tr_Queued() int64 {
	return this.tr_queued
}

func (this *Message) Tr_Running() int64 {
	return this.tr_running
}

func (this *Message) Tr_End() int64 {
	return this.tr_end
}

func (this *Message) Tr_Timeout() int64 {
	return this.tr_timeout
}

func (this *Message) Tr_Replica() byte {
	return this.tr_replica
}

func (this *Message) Tr_IsReplica() bool {
	return this.tr_isReplica
}

//Setters

func (this *Message) SetSource(source string) {
	this.source = source
}

func (this *Message) SetVnet(vnet string) {
	this.vnet = vnet
}

func (this *Message) SetDestination(destination string) {
	this.destination = destination
}

func (this *Message) SetServiceName(serviceName string) {
	this.serviceName = serviceName
}

func (this *Message) SetServiceArea(serviceArea byte) {
	this.serviceArea = serviceArea
}

func (this *Message) SetSequence(sequence uint32) {
	this.sequence = sequence
}

func (this *Message) SetPriority(priority Priority) {
	this.priority = priority
}

func (this *Message) SetMulticastMode(multicastMode MulticastMode) {
	this.multicastMode = multicastMode
}

func (this *Message) SetAction(action Action) {
	this.action = action
}

func (this *Message) SetTimeout(timeout uint16) {
	this.timeout = timeout
}

func (this *Message) SetRequestReply(request, reply bool) {
	this.request = request
	this.reply = reply
}

func (this *Message) SetFailMessage(failMessage string) {
	this.failMessage = failMessage
}

func (this *Message) SetAAAId(aaaId string) {
	this.aaaId = aaaId
}

func (this *Message) SetData(data []byte) {
	this.data = data
}

func (this *Message) SetTr_State(trstate TransactionState) {
	this.tr_state = trstate
	switch trstate {
	case Created:
		this.tr_created = time.Now().Unix()
	case Queued:
		this.tr_queued = time.Now().Unix()
	case Running:
		this.tr_running = time.Now().Unix()
	case Failed:
		fallthrough
	case Committed:
		this.tr_end = time.Now().Unix()
	}
}

func (this *Message) SetTr_Id(trid string) {
	this.tr_id = trid
}

func (this *Message) SetTr_ErrMsg(errMsg string) {
	this.tr_errMsg = errMsg
}

func (this *Message) SetTr_Timeout(timeout int64) {
	this.tr_timeout = timeout
}

func (this *Message) SetTr_Replica(replica byte) {
	this.tr_replica = replica
}

func (this *Message) SetTr_IsReplica(isReplica bool) {
	this.tr_isReplica = isReplica
}
