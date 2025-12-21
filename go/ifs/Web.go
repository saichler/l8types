/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
