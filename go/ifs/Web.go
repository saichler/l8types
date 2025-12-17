package ifs

import (
	"google.golang.org/protobuf/proto"
	"net/http"

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
	Vnet() uint32
	ServiceName() string
	ServiceArea() byte
	Protos(string, Action) (proto.Message, proto.Message, error)
	AddEndpoint(proto.Message, Action, proto.Message)
	Serialize() *l8web.L8WebService
	DeSerialize(*l8web.L8WebService, IRegistry) error
	Plugin() string
}

type IPlugin interface {
	Install(IVNic) error
}

type IWebProxy interface {
	RegisterHandlers(mux *http.ServeMux)
	ProxyRequest(w http.ResponseWriter, r *http.Request) error
	SetValidator(BearerValidator)
}

type BearerValidator interface {
	ValidateBearerToken(r *http.Request) error
}
