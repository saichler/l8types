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
	"reflect"

	"github.com/saichler/l8types/go/types/l8api"
)

/*
IRegistry - Interface for encapsulating a Type registry service so a struct instance
can be instantiated based on the type name.
*/
type IRegistry interface {
	//Register - receive as an input an instance and register, extract the Type and register it.
	Register(interface{}) (bool, error)
	//RegisterType - receive a reflect.Type instance and register it.
	RegisterType(reflect.Type) (bool, error)
	//Info Retrieve the registered entry for this type.
	Info(string) (IInfo, error)
	//Register Enum string to int32 values
	RegisterEnums(map[string]int32)
	//Get int32 value of an enum
	Enum(string) int32
	//Unregister an instance
	UnRegister(string) (bool, error)
	NewOf(interface{}) interface{}
	TypeList() *l8api.L8TypeList
}

type IInfo interface {
	Type() reflect.Type
	Name() string
	Serializer(SerializerMode) ISerializer
	AddSerializer(ISerializer)
	NewInstance() (interface{}, error)
}
