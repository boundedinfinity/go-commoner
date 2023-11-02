package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func New() *Marshaler {
	return &Marshaler{
		types: make(map[string]marshalerContext),
	}
}

type MarshalDiscriminator interface {
	Discriminator() string
}

type marshalerContext struct {
	typ reflect.Type
	gen func() any
}

type Marshaler struct {
	discriminator MarshalDiscriminator
	types         map[string]marshalerContext
}

func (t *Marshaler) Interface(discriminator MarshalDiscriminator) {
	t.discriminator = discriminator
}

func (t *Marshaler) Register(item any, gen func() any) {
	ctx := marshalerContext{
		typ: reflect.TypeOf(item),
		gen: gen,
	}

	t.types[ctx.typ.Name()] = ctx
}

func (t Marshaler) UnmarshalInterface(data []byte) (any, error) {
	var err error

	if err = json.Unmarshal(data, t.discriminator); err != nil {
		return nil, err
	}

	ctx, ok := t.types[t.discriminator.Discriminator()]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", t.discriminator.Discriminator())
	}

	new := reflect.New(ctx.typ)
	ptr := new.Elem().Interface()
	val := new.Elem().Interface()

	if err = json.Unmarshal(data, ptr); err != nil {
		return nil, err
	}

	return val, err
}

func x() {
	var x string
	reflect.TypeOf(x).Name()
}
