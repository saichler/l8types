package ifs

type IWebServer interface {
	RegisterWebService(IWebService)
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
}

type IPlugin interface {
	Install(IVNic) error
}
