package marshaler

import (
	"encoding/json"

	"reflect"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// Reference
// https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0

var (
	ErrTypedMarshalerTypeMissing            = errorer.New("type missing")
	ErrTypedMarshalerTypeNotRegistered      = errorer.New("type not registered")
	ErrTypedMarshalerTypeNotRegisteredv     = ErrTypedMarshalerTypeNotRegistered.ValueFn()
	ErrTypedMarshalerTypeAlreadyRegistered  = errorer.New("type already registered")
	ErrTypedMarshalerTypeAlreadyRegisteredv = ErrTypedMarshalerTypeAlreadyRegistered.ValueFn()
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
	typ  reflect.Type
}

type TypedMarshaler struct {
	types map[string]typedInfo
}

func (t *TypedMarshaler) Register(namers ...TypeNamer) error {
	for _, elem := range namers {
		if err := t.register(elem); err != nil {
			return err
		}
	}

	return nil
}

func (t *TypedMarshaler) register(namer TypeNamer) error {
	name := namer.TypeName()

	if stringer.IsEmpty(name) {
		return ErrTypedMarshalerTypeMissing
	}

	if _, ok := t.types[name]; ok {
		return ErrTypedMarshalerTypeAlreadyRegisteredv(name)
	}

	t.types[name] = typedInfo{
		name: name,
		typ:  reflect.TypeOf(namer),
	}

	return nil
}

func (t TypedMarshaler) Marshal(namer TypeNamer) ([]byte, error) {
	name := namer.TypeName()

	if stringer.IsEmpty(name) {
		return nil, ErrTypedMarshalerTypeMissing
	}

	if _, ok := t.types[name]; !ok {
		return nil, ErrTypedMarshalerTypeNotRegisteredv(name)
	}

	typed := typedMarshaler{
		Type:  name,
		Value: namer,
	}

	return json.Marshal(typed)
}

func (t TypedMarshaler) Unmarshal(data []byte) (any, error) {
	var err error
	var discriminator typedUnmarshaler

	if err = json.Unmarshal(data, &discriminator); err != nil {
		return nil, err
	}

	if stringer.IsEmpty(discriminator.Type) {
		return nil, ErrTypedMarshalerTypeMissing
	}

	info, ok := t.types[discriminator.Type]

	if !ok {
		return nil, ErrTypedMarshalerTypeNotRegisteredv(discriminator.Type)
	}

	rf := reflect.New(info.typ)

	if err = json.Unmarshal(discriminator.Value, rf.Interface()); err != nil {
		return nil, err
	}

	return rf.Elem().Interface(), err
}
