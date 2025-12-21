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
)

const (
	Deleted_Entry = "__DD__"
)

type IElements interface {
	Elements() []interface{}
	Keys() []interface{}
	Errors() []error
	Element() interface{}
	Query(IResources) (IQuery, error)
	Key() interface{}
	Error() error
	Serialize() ([]byte, error)
	Deserialize([]byte, IRegistry) error
	Notification() bool
	Append(IElements)
	AsList(IRegistry) (interface{}, error)
	IsFilterMode() bool
	IsReplica() bool
	Replica() byte
}

type IQuery interface {
	RootType() *l8reflect.L8Node
	Properties() []IProperty
	Criteria() IExpression
	KeyOf() string
	Match(interface{}) bool
	Page() int32
	Limit() int32
	SortBy() string
	SortByValue(interface{}) interface{}
	MatchCase() bool
	Descending() bool
	MapReduce() bool
	Text() string
	Hash() string
	ValueForParameter(string) string
}

type IProperty interface {
	PropertyId() (string, error)
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) (interface{}, interface{}, error)
	Node() *l8reflect.L8Node
	Parent() IProperty
	IsString() bool
	Resources() IResources
}

type IExpression interface {
	Condition() ICondition
	Operator() string
	Next() IExpression
	Child() IExpression
}

type ICondition interface {
	Comparator() IComparator
	Operator() string
	Next() ICondition
}

type IComparator interface {
	Left() string
	LeftProperty() IProperty
	Right() string
	RightProperty() IProperty
	Operator() string
}
