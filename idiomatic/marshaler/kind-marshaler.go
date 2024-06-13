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
	ErrKindMarshalerTypeMissing            = errorer.New("type missing")
	ErrKindMarshalerTypeNotRegistered      = errorer.New("type not registered")
	ErrKindMarshalerTypeNotRegisteredv     = ErrKindMarshalerTypeNotRegistered.ValueFn()
	ErrKindMarshalerTypeAlreadyRegistered  = errorer.New("type already registered")
	ErrKindMarshalerTypeAlreadyRegisteredv = ErrKindMarshalerTypeAlreadyRegistered.ValueFn()
)

func NewKind[D any](val D, fn func(D) string) *KindMarshaler[D] {
	return &KindMarshaler[D]{
		discriminatorType: reflect.TypeOf(val),
		discriminatorFn:   fn,
		lookup:            make(map[string]kindInfo),
	}
}

type KindMarshaler[D any] struct {
	discriminatorType reflect.Type
	discriminatorFn   func(D) string
	lookup            map[string]kindInfo
	Formatted         bool
}

type kindInfo struct {
	discriminator string
	typ           reflect.Type
}

func (t *KindMarshaler[D]) RegisterNamer(namer TypeNamer) error {
	return t.RegisterValue(namer.TypeName(), namer)
}

func (t *KindMarshaler[D]) RegisterValue(discriminator string, val any) error {
	return t.RegisterType(discriminator, reflect.TypeOf(val))
}

func (t *KindMarshaler[D]) RegisterType(discriminator string, typ reflect.Type) error {
	if stringer.IsEmpty(discriminator) {
		return ErrKindMarshalerTypeMissing
	}

	if _, ok := t.lookup[discriminator]; ok {
		return ErrKindMarshalerTypeAlreadyRegisteredv(discriminator)
	}

	t.lookup[discriminator] = kindInfo{
		discriminator: discriminator,
		typ:           typ,
	}

	return nil
}

func (t KindMarshaler[D]) Marshal(val any) ([]byte, error) {
	if t.Formatted {
		return json.MarshalIndent(val, "", "    ")
	} else {
		return json.Marshal(val)
	}
}

func (t KindMarshaler[D]) Unmarshal(data []byte) (any, error) {
	dv := reflect.New(t.discriminatorType)

	if err := json.Unmarshal(data, dv.Interface()); err != nil {
		return nil, err
	}

	discriminator := t.discriminatorFn(dv.Elem().Interface().(D))

	if stringer.IsEmpty(discriminator) {
		return nil, ErrKindMarshalerTypeMissing
	}

	info, ok := t.lookup[discriminator]

	if !ok {
		return nil, ErrKindMarshalerTypeNotRegisteredv(discriminator)
	}

	vv := reflect.New(info.typ)

	if err := json.Unmarshal(data, vv.Interface()); err != nil {
		return nil, err
	}

	return vv.Elem().Interface(), nil
}
