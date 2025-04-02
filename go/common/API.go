package common

import "github.com/saichler/types/go/types"

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
}

type IQuery interface {
	RootType() *types.RNode
	Properties() []IProperty
}

type IProperty interface {
	PropertyId() (string, error)
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) (interface{}, interface{}, error)
}
