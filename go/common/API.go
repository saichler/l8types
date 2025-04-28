package common

import "github.com/saichler/types/go/types"

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
	ReplicasRequest() bool
}

type IQuery interface {
	RootType() *types.RNode
	Properties() []IProperty
	Criteria() IExpression
	KeyOf() string
}

type IProperty interface {
	PropertyId() (string, error)
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) (interface{}, interface{}, error)
	Node() *types.RNode
	Parent() IProperty
	IsString() bool
	Introspector() IIntrospector
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
