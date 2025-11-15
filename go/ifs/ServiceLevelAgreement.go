package ifs

type ServiceLevelAgreement struct {
	serviceHandlerInstance IServiceHandler
	serviceName            string
	serviceArea            byte
	stateful               bool
	callback               IServiceCallback

	serviceItem     interface{}
	serviceItemList interface{}
	initItems       []interface{}
	primaryKeys     []string
	store           IStorage

	voter            bool
	transactional    bool
	replication      bool
	replicationCount int

	webService IWebService
	args       []interface{}

	metadataFunc map[string]func(interface{}) (bool, string)
}

func NewServiceLevelAgreement(serviceHandlerInstance IServiceHandler, serviceName string, serviceArea byte, stateful bool, callback IServiceCallback) *ServiceLevelAgreement {
	return &ServiceLevelAgreement{serviceHandlerInstance: serviceHandlerInstance, serviceName: serviceName, serviceArea: serviceArea, callback: callback, stateful: stateful}
}

// Getters and Setters for attributes not in constructor
func (this *ServiceLevelAgreement) ServiceName() string {
	return this.serviceName
}

func (this *ServiceLevelAgreement) ServiceArea() byte {
	return this.serviceArea
}

func (this *ServiceLevelAgreement) Stateful() bool {
	return this.stateful
}

func (this *ServiceLevelAgreement) ServiceHandlerInstance() IServiceHandler {
	return this.serviceHandlerInstance
}

func (this *ServiceLevelAgreement) Callback() IServiceCallback {
	return this.callback
}

func (this *ServiceLevelAgreement) ServiceItem() interface{} {
	return this.serviceItem
}

func (this *ServiceLevelAgreement) SetServiceItem(serviceItem interface{}) {
	this.serviceItem = serviceItem
}

func (this *ServiceLevelAgreement) ServiceItemList() interface{} {
	return this.serviceItemList
}

func (this *ServiceLevelAgreement) SetServiceItemList(serviceItemList interface{}) {
	this.serviceItemList = serviceItemList
}

func (this *ServiceLevelAgreement) InitItems() []interface{} {
	return this.initItems
}

func (this *ServiceLevelAgreement) SetInitItems(initItems []interface{}) {
	this.initItems = initItems
}

func (this *ServiceLevelAgreement) PrimaryKeys() []string {
	return this.primaryKeys
}

func (this *ServiceLevelAgreement) SetPrimaryKeys(primaryKeys ...string) {
	this.primaryKeys = primaryKeys
}

func (this *ServiceLevelAgreement) Store() IStorage {
	return this.store
}

func (this *ServiceLevelAgreement) SetStore(store IStorage) {
	this.store = store
}

func (this *ServiceLevelAgreement) Voter() bool {
	return this.voter
}

func (this *ServiceLevelAgreement) SetVoter(voter bool) {
	this.voter = voter
}

func (this *ServiceLevelAgreement) Transactional() bool {
	return this.transactional
}

func (this *ServiceLevelAgreement) SetTransactional(transactional bool) {
	this.transactional = transactional
}

func (this *ServiceLevelAgreement) Replication() bool {
	return this.replication
}

func (this *ServiceLevelAgreement) SetReplication(replication bool) {
	this.replication = replication
}

func (this *ServiceLevelAgreement) ReplicationCount() int {
	return this.replicationCount
}

func (this *ServiceLevelAgreement) SetReplicationCount(replicationCount int) {
	this.replicationCount = replicationCount
}

func (this *ServiceLevelAgreement) WebService() IWebService {
	return this.webService
}

func (this *ServiceLevelAgreement) SetWebService(webService IWebService) {
	this.webService = webService
}

func (this *ServiceLevelAgreement) Args() []interface{} {
	return this.args
}

func (this *ServiceLevelAgreement) SetArgs(args ...interface{}) {
	this.args = args
}

func (this *ServiceLevelAgreement) AddMetadataFunc(name string, f func(interface{}) (bool, string)) {
	if this.metadataFunc == nil {
		this.metadataFunc = make(map[string]func(interface{}) (bool, string))
	}
	this.metadataFunc[name] = f
}

func (this *ServiceLevelAgreement) MetadataFunc() map[string]func(interface{}) (bool, string) {
	return this.metadataFunc
}

type IServiceCallback interface {
	BeforePost(IElements, IVNic) IElements
	AfterPost(IElements, IVNic) IElements
	BeforePut(IElements, IVNic) IElements
	AfterPut(IElements, IVNic) IElements
	BeforePatch(IElements, IVNic) IElements
	AfterPatch(IElements, IVNic) IElements
	BeforeDelete(IElements, IVNic) IElements
	AfterDelete(IElements, IVNic) IElements
	BeforeGet(IElements, IVNic) IElements
	AfterGet(IElements, IVNic) IElements
}
