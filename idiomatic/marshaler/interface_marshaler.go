package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewInterfaceMarshaler(discriminator InterfaceUnmarshalDiscriminator) *InterfaceMarshaler {
	return &InterfaceMarshaler{
		discriminator: discriminator,
		types:         make(map[string]reflect.Type),
	}
}

type InterfaceUnmarshalDiscriminator interface {
	Discriminator() string
}

type InterfaceMarshaler struct {
	discriminator InterfaceUnmarshalDiscriminator
	types         map[string]reflect.Type
}

func (t *InterfaceMarshaler) Register(items ...any) {
	for _, item := range items {
		t.register(item)
	}
}

func (t *InterfaceMarshaler) register(item any) {
	typ := reflect.TypeOf(item)
	t.types[typ.Name()] = typ
}

func (t InterfaceMarshaler) Marshal(item any) ([]byte, error) {
	return json.Marshal(item)
}

func (t InterfaceMarshaler) Unmarshal(data []byte) (any, error) {
	var err error

	if err = json.Unmarshal(data, t.discriminator); err != nil {
		return nil, err
	}

	typ, ok := t.types[t.discriminator.Discriminator()]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", t.discriminator.Discriminator())
	}

	ptr := reflect.New(typ)

	if err = json.Unmarshal(data, ptr.Interface()); err != nil {
		return nil, err
	}

	return ptr.Elem().Interface(), err
}
