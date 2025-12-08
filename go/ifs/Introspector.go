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

	PrimaryKeyValue(interface{}) (string, error)
	UniqueKeyValue(interface{}) (string, error)
}
