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

import "reflect"

type SerializerMode int

const (
	BINARY SerializerMode = 1
	JSON   SerializerMode = 2
	STRING SerializerMode = 3
)

type ISerializer interface {
	Mode() SerializerMode
	Marshal(interface{}, IResources) ([]byte, error)
	Unmarshal([]byte, IResources) (interface{}, error)
}

func IsNil(any interface{}) bool {
	if any == nil {
		return true
	}
	v := reflect.ValueOf(any)
	isNil := v.IsNil()
	if !isNil {
		if v.Kind() == reflect.Func {
			panic("Trying to check nil on a function!")
		}
	}
	return isNil
}

type IStorage interface {
	Put(string, interface{}) error
	Get(string) (interface{}, error)
	Delete(string) (interface{}, error)
	Collect(f func(interface{}) (bool, interface{})) map[string]interface{}
	CacheEnabled() bool
}
