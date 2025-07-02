package ifs

type Message struct {
	source      string
	vnet        string
	destination string
	serviceName string
	serviceArea byte
	priority    Priority
	action      Action
	tr_state    TransactionState

	aaaId       string
	sequence    uint32
	timeout     uint16
	request     bool
	reply       bool
	failMessage string
	data        string

	tr_id        string
	tr_errMsg    string
	tr_startTime int64
}

func NewMessage(destination, serviceName string, serviceArea byte,
	priority Priority, action Action, source, vnet string, data string,
	isRequest, isReply bool, msgNum uint32,
	tr_state TransactionState, tr_id, tr_errMsg string, tr_start int64) *Message {

	return &Message{
		destination:  destination,
		serviceName:  serviceName,
		serviceArea:  serviceArea,
		priority:     priority,
		action:       action,
		source:       source,
		vnet:         vnet,
		data:         data,
		request:      isRequest,
		reply:        isReply,
		sequence:     msgNum,
		tr_state:     tr_state,
		tr_id:        tr_id,
		tr_errMsg:    tr_errMsg,
		tr_startTime: tr_start,
	}
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

func (this *Message) Data() string {
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

func (this *Message) Tr_StartTime() int64 {
	return this.tr_startTime
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

func (this *Message) SetData(data string) {
	this.data = data
}

func (this *Message) SetTr_State(trstate TransactionState) {
	this.tr_state = trstate
}

func (this *Message) SetTr_Id(trid string) {
	this.tr_id = trid
}

func (this *Message) SetTr_ErrMsg(errMsg string) {
	this.tr_errMsg = errMsg
}

func (this *Message) SetTr_StartTime(trStartTime int64) {
	this.tr_startTime = trStartTime
}
