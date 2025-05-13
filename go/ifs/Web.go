package ifs

import "github.com/saichler/l8types/go/types"

type IWebServer interface {
	RegisterWebService(IWebService, IVNic)
	Start() error
}

type IWebService interface {
	ServiceName() string
	ServiceArea() uint16

	PostBody() string
	PostResp() string

	PutBody() string
	PutResp() string

	PatchBody() string
	PatchResp() string

	DeleteBody() string
	DeleteResp() string

	GetBody() string
	GetResp() string

	Serialize() *types.WebService
	DeSerialize(*types.WebService)

	Plugin() string
}

type IPlugin interface {
	Install(IVNic) error
}
