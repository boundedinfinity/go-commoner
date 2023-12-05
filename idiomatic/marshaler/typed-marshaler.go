package marshaler

import (
	"encoding/json"
	"fmt"

	greflect "reflect"
)

type TypeNamer interface {
	TypeName() string
}

func NewTyped() *TypedMarshaler {
	return &TypedMarshaler{
		types: make(map[string]typedInfo),
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

type typedInfo struct {
	name string
	typ  greflect.Type
}

type TypedMarshaler struct {
	types map[string]typedInfo
}

func (t *TypedMarshaler) Register(namers ...TypeNamer) {
	for _, item := range namers {
		t.register(item)
	}
}

func (t *TypedMarshaler) register(namer TypeNamer) {
	name := namer.TypeName()
	t.types[name] = typedInfo{
		name: name,
		typ:  greflect.TypeOf(namer),
	}
}

func (t TypedMarshaler) Marshal(namer TypeNamer) ([]byte, error) {
	name := namer.TypeName()

	if _, ok := t.types[name]; !ok {
		return nil, fmt.Errorf("no type found for %v", name)
	}

	typed := typedMarshaler{
		Type:  name,
		Value: namer,
	}

	return json.Marshal(typed)
}

func (t TypedMarshaler) Unmarshal(data []byte) (any, error) {
	var err error
	var typed typedUnmarshaler

	if err = json.Unmarshal(data, &typed); err != nil {
		return nil, err
	}

	info, ok := t.types[typed.Type]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", typed.Type)
	}

	rf := greflect.New(info.typ)

	if err = json.Unmarshal(typed.Value, rf.Interface()); err != nil {
		return nil, err
	}

	return rf.Elem().Interface(), err
}
