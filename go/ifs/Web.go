package ifs

import "google.golang.org/protobuf/proto"

type IWebServer interface {
	NewWebServiceHandler(string, uint16, IVNic) IWebServiceHandler
}

type IWebServiceHandler interface {
	ServiceName() string
	ServiceArea() uint16
	PostBodyResponse(body proto.Message, resp proto.Message)
	PutBodyResponse(body proto.Message, resp proto.Message)
	PatchBodyResponse(body proto.Message, resp proto.Message)
	DeleteBodyResponse(body proto.Message, resp proto.Message)
	GetBodyResponse(body proto.Message, resp proto.Message)
}

type IServicePlugin interface {
	InstallRegistry(IVNic) error
	InstallServices(IVNic) error
}
