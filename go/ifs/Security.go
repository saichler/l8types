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
	"errors"
	"fmt"
	"net"
	"plugin"

	"github.com/saichler/l8types/go/types/l8sysconfig"
)

type ISecurityProvider interface {
	Authenticate(string, string) (string, bool, bool, error)
	ValidateToken(string) (string, bool)

	Message(string) (*Message, error)

	CanDial(string, uint32) (net.Conn, error)
	CanAccept(net.Conn) error
	ValidateConnection(net.Conn, *l8sysconfig.L8SysConfig) error

	Encrypt([]byte) (string, error)
	Decrypt(string) ([]byte, error)

	CanDoAction(Action, IElements, string, string, ...string) error
	ScopeView(IElements, string, string, ...string) IElements

	TFASetup(string, IVNic) (string, []byte, error)
	TFAVerify(string, string, string, IVNic) error

	Captcha() []byte
	Register(string, string, string, IVNic) error

	Credential(string, string, IResources) (string, string, string, string, error)
}

type ISecurityProviderLoader interface {
	LoadSecurityProvider(...interface{}) (ISecurityProvider, error)
}

type ISecurityProviderActivate interface {
	Activate(IVNic)
}

func LoadSecurityProvider(args ...interface{}) (ISecurityProvider, error) {
	loaderFile, err := plugin.Open("/var/loader.so")
	if err != nil {
		fmt.Println("Failed to load security provider #1: ", err.Error())
		return nil, errors.New("Failed to load security provider #1: " + err.Error())
	}
	loader, err := loaderFile.Lookup("Loader")
	if err != nil {
		fmt.Println("Failed to load security provider #2: ", err.Error())
		return nil, errors.New("Failed to load security provider #2: " + err.Error())
	}
	if loader == nil {
		fmt.Println("Failed to load security provider #3: Nil Loader")
		return nil, errors.New("Failed to load security provider #3: Nil Loader")
	}
	loaderInterface := *loader.(*ISecurityProviderLoader)
	securityLoader := loaderInterface.(ISecurityProviderLoader).(ISecurityProviderLoader)
	return securityLoader.LoadSecurityProvider(args...)
}
