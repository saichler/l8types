package ifs

import (
	"github.com/saichler/l8types/go/types/l8web"
)

const (
	WebService = "WebService"
)

type IWebServer interface {
	RegisterWebService(IWebService, IVNic)
	Start() error
	Stop()
}

type IWebService interface {
	ServiceName() string
	ServiceArea() byte

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

	Serialize() *l8web.L8WebService
	DeSerialize(*l8web.L8WebService)

	Plugin() string
}

type IPlugin interface {
	Install(IVNic) error
}
