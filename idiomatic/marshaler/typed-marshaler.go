package marshaler

import (
	"encoding/json"
	"fmt"

	greflect "reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
)

func NewTyped() *TypedMarshaler {
	return &TypedMarshaler{
		types: make(map[string]greflect.Type),
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
	types map[string]greflect.Type
}

func (t *TypedMarshaler) Register(items ...any) {
	for _, item := range items {
		t.register(item)
	}
}

func (t *TypedMarshaler) register(item any) {
	name := reflecter.Instances.QualifiedName(item)
	typ := greflect.TypeOf(item)
	t.types[name] = typ
}

func (t TypedMarshaler) Marshal(item any) ([]byte, error) {
	name := reflecter.Instances.QualifiedName(item)

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
	var typed typedUnmarshaler

	if err = json.Unmarshal(data, &typed); err != nil {
		return nil, err
	}

	typ, ok := t.types[typed.Type]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", typed.Type)
	}

	rf := greflect.New(typ)

	if err = json.Unmarshal(typed.Value, rf.Interface()); err != nil {
		return nil, err
	}

	return rf.Elem().Interface(), err
}
