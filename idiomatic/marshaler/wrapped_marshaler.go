package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewWrapped() *WrappedMarshaler {
	return &WrappedMarshaler{
		types: make(map[string]reflect.Type),
	}
}

type marshalerWrapper struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type unmarshalerWrapper struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type WrappedMarshaler struct {
	types map[string]reflect.Type
}

func (t *WrappedMarshaler) Register(items ...any) {
	for _, item := range items {
		t.register(item)
	}
}

func (t *WrappedMarshaler) register(item any) {
	typ := reflect.TypeOf(item)

	t.types[typ.Name()] = typ
}

func (t WrappedMarshaler) Marshal(item any) ([]byte, error) {
	name := reflect.TypeOf(item).Name()

	if _, ok := t.types[name]; !ok {
		return nil, fmt.Errorf("no type found for %v", name)
	}

	wrapped := marshalerWrapper{
		Type:  name,
		Value: item,
	}

	return json.Marshal(wrapped)
}

func (t WrappedMarshaler) Unmarshal(data []byte) (any, error) {
	var err error
	var wrapped unmarshalerWrapper

	if err = json.Unmarshal(data, &wrapped); err != nil {
		return nil, err
	}

	typ, ok := t.types[wrapped.Type]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", wrapped.Type)
	}

	rf := reflect.New(typ)

	// fmt.Printf("vp.Type(): %v\n", rf.Type())
	// fmt.Printf("vp.Elem(): %v\n", rf.Elem())
	// fmt.Printf("vp.Interface(): %v\n", rf.Interface())
	// fmt.Printf("vp.Type().Elem(): %v\n", rf.Type().Elem())
	// fmt.Printf("vp.Elem().Interface(): %v\n", rf.Elem().Interface())

	if err = json.Unmarshal(wrapped.Value, rf.Interface()); err != nil {
		return nil, err
	}

	return rf.Elem().Interface(), err
}
