package ifs

type ServiceLevelAgreement struct {
	serviceHandlerInstance IServiceHandler
	serviceName            string
	serviceArea            byte
	stateful               bool
	callback               IServiceCallback

	serviceItem      interface{}
	serviceItemList  interface{}
	initItems        []interface{}
	primaryKeys      []string
	store            IStorage
	voter            bool
	transactional    bool
	replication      bool
	replicationCount int
	webServiceDef    IWebService
}

func NewServiceLevelAgreement(serviceHandlerInstance IServiceHandler, serviceName string, serviceArea byte, stateful bool, callback IServiceCallback) *ServiceLevelAgreement {
	return &ServiceLevelAgreement{serviceHandlerInstance: serviceHandlerInstance, serviceName: serviceName, serviceArea: serviceArea, callback: callback, stateful: stateful}
}

// Getters and Setters for attributes not in constructor

func (s *ServiceLevelAgreement) ServiceItem() interface{} {
	return s.serviceItem
}

func (s *ServiceLevelAgreement) SetServiceItem(serviceItem interface{}) {
	s.serviceItem = serviceItem
}

func (s *ServiceLevelAgreement) ServiceItemList() interface{} {
	return s.serviceItemList
}

func (s *ServiceLevelAgreement) SetServiceItemList(serviceItemList interface{}) {
	s.serviceItemList = serviceItemList
}

func (s *ServiceLevelAgreement) InitItems() []interface{} {
	return s.initItems
}

func (s *ServiceLevelAgreement) SetInitItems(initItems []interface{}) {
	s.initItems = initItems
}

func (s *ServiceLevelAgreement) PrimaryKeys() []string {
	return s.primaryKeys
}

func (s *ServiceLevelAgreement) SetPrimaryKeys(primaryKeys []string) {
	s.primaryKeys = primaryKeys
}

func (s *ServiceLevelAgreement) Store() IStorage {
	return s.store
}

func (s *ServiceLevelAgreement) SetStore(store IStorage) {
	s.store = store
}

func (s *ServiceLevelAgreement) Voter() bool {
	return s.voter
}

func (s *ServiceLevelAgreement) SetVoter(voter bool) {
	s.voter = voter
}

func (s *ServiceLevelAgreement) Transactional() bool {
	return s.transactional
}

func (s *ServiceLevelAgreement) SetTransactional(transactional bool) {
	s.transactional = transactional
}

func (s *ServiceLevelAgreement) Replication() bool {
	return s.replication
}

func (s *ServiceLevelAgreement) SetReplication(replication bool) {
	s.replication = replication
}

func (s *ServiceLevelAgreement) ReplicationCount() int {
	return s.replicationCount
}

func (s *ServiceLevelAgreement) SetReplicationCount(replicationCount int) {
	s.replicationCount = replicationCount
}

func (s *ServiceLevelAgreement) WebServiceDef() IWebService {
	return s.webServiceDef
}

func (s *ServiceLevelAgreement) SetWebServiceDef(webServiceDef IWebService) {
	s.webServiceDef = webServiceDef
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
