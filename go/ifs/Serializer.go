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
