package ifs

type IWebServer interface {
	RegisterWebService(IWebService)
	Start() error
}

type IWebService interface {
	ServiceName() string
	ServiceArea() uint16

	PostBodyType() string
	PostRespType() string

	PutBodyType() string
	PutRespType() string

	PatchBodyType() string
	PatchRespType() string

	DeleteBodyType() string
	DeleteRespType() string

	GetBodyType() string
	GetRespType() string

	Plugin() IPlugin
}

type IPlugin interface {
	Install(IVNic) error
}
