package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewInterface() *KindMarshaler {
	return &KindMarshaler{
		types: make(map[string]kindContext),
	}
}

type kindContext struct {
	kind KindUnmarshal
	typ  reflect.Type
}

type KindMarshaler struct {
	types map[string]kindContext
}

type KindUnmarshal interface {
	Kind() string
}

func (t *KindMarshaler) Register(kind KindUnmarshal) {
	switch typ := kind.(type) {
	default:
		if _, ok := t.types[kind.Kind()]; ok {
			fmt.Printf("already registered %v", kind)
		} else {
			t.types[kind.Kind()] = kindContext{
				kind: kind,
				typ:  reflect.TypeOf(typ),
			}
		}
	}
}

func (t KindMarshaler) Marshal(item any) ([]byte, error) {
	return json.Marshal(item)
}

func (t KindMarshaler) Unmarshal(data []byte) (any, error) {
	var err error

	// if err = json.Unmarshal(data, t.kind); err != nil {
	// 	return nil, err
	// }

	// typ, ok := t.types[t.kind.Kind()]

	// if !ok {
	// 	return nil, fmt.Errorf("no type found for %v", t.kind.Kind())
	// }

	// ptr := reflect.New(typ)

	// if err = json.Unmarshal(data, ptr.Interface()); err != nil {
	// 	return nil, err
	// }

	// return ptr.Elem().Interface(), err
	return nil, err
}
