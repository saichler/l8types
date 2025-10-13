package ifs

type ServiceConfig struct {
	ServiceName       string
	ServiceArea       byte
	ServiceItem       interface{}
	ServiceItemList   interface{}
	InitItems         []interface{}
	PrimaryKey        []string
	Store             IStorage
	SendNotifications bool
	Transaction       bool
	Replication       bool
	ReplicationCount  int
	WebServiceDef     IWebService
}
