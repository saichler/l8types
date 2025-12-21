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
	"github.com/saichler/l8types/go/types/l8reflect"

	"reflect"
)

type IIntrospector interface {
	Inspect(interface{}) (*l8reflect.L8Node, error)
	Node(string) (*l8reflect.L8Node, bool)
	NodeByType(p reflect.Type) (*l8reflect.L8Node, bool)
	NodeByTypeName(string) (*l8reflect.L8Node, bool)
	NodeByValue(interface{}) (*l8reflect.L8Node, bool)
	Nodes(bool, bool) []*l8reflect.L8Node
	Registry() IRegistry
	Kind(*l8reflect.L8Node) reflect.Kind
	Clone(interface{}) interface{}
	TableView(string) (*l8reflect.L8TableView, bool)
	TableViews() []*l8reflect.L8TableView
	Clean(string)
	Decorators() IDecorators
}

type IDecorators interface {
	AddPrimaryKeyDecorator(interface{}, ...string) error
	AddUniqueKeyDecorator(interface{}, ...string) error
	AddNonUniqueKeyDecorator(interface{}, ...string) error

	AddAlwayOverwriteDecorator(string) error
	AddNoNestedInspection(interface{}) error

	PrimaryKeyDecoratorValue(interface{}) (string, *l8reflect.L8Node, error)
	PrimaryKeyDecoratorFromValue(*l8reflect.L8Node, reflect.Value) (string, *l8reflect.L8Node, error)
	UniqueKeyDecoratorValue(interface{}) (string, *l8reflect.L8Node, error)
	NonUniqueKeyDecoratorValue(interface{}) (string, *l8reflect.L8Node, error)
	NoNestedInspection(interface{}) bool
	AlwaysFullDecorator(interface{}) bool

	NodeFor(interface{}) (*l8reflect.L8Node, reflect.Value, error)
	BoolDecoratorValueFor(interface{}, l8reflect.L8DecoratorType) bool
	BoolDecoratorValueForNode(*l8reflect.L8Node, l8reflect.L8DecoratorType) bool
	Fields(*l8reflect.L8Node, l8reflect.L8DecoratorType) ([]string, error)
	KeyForValue([]string, reflect.Value, string, bool) (string, error)
}
