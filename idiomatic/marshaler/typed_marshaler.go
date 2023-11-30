package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewTyped() *TypedMarshaler {
	return &TypedMarshaler{
		types: make(map[string]reflect.Type),
	}
}

type typedMarshaler struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type typedUnmarshaler struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type TypedMarshaler struct {
	types map[string]reflect.Type
}

func (t *TypedMarshaler) Register(items ...any) {
	for _, item := range items {
		t.register(item)
	}
}

func (t *TypedMarshaler) register(item any) {
	typ := reflect.TypeOf(item)

	t.types[typ.Name()] = typ
}

func (t TypedMarshaler) Marshal(item any) ([]byte, error) {
	name := reflect.TypeOf(item).Name()

	if _, ok := t.types[name]; !ok {
		return nil, fmt.Errorf("no type found for %v", name)
	}

	typed := typedMarshaler{
		Type:  name,
		Value: item,
	}

	return json.Marshal(typed)
}

func (t TypedMarshaler) Unmarshal(data []byte) (any, error) {
	var err error
	var wrapped typedUnmarshaler

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
