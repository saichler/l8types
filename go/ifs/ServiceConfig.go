package ifs

type ServiceConfig struct {
	ServiceName      string
	ServiceArea      byte
	ServiceItem      interface{}
	ServiceItemList  interface{}
	InitItems        []interface{}
	PrimaryKey       []string
	Store            IStorage
	Voter            bool
	Transaction      bool
	Replication      bool
	ReplicationCount int
	WebServiceDef    IWebService
	Callback         IServiceCallback
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
