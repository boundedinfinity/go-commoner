package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewKind[K comparable, D any]() *KindMarshaler[K, D] {
	return &KindMarshaler[K, D]{
		lookup: make(map[K]kindInfo[K]),
	}
}

type KindMarshaler[K comparable, D any] struct {
	discriminatorType reflect.Type
	discriminatorFn   func(D) K
	lookup            map[K]kindInfo[K]
	Formatted         bool
}

type kindInfo[K comparable] struct {
	discriminator K
	typ           reflect.Type
}

func (t *KindMarshaler[K, D]) RegisterDescriminator(val D, fn func(D) K) {
	t.discriminatorType = reflect.TypeOf(val)
	t.discriminatorFn = fn
}

func (t *KindMarshaler[K, D]) RegisterValue(discriminator K, val any) {
	t.RegisterType(discriminator, reflect.TypeOf(val))
}

func (t *KindMarshaler[K, D]) RegisterType(discriminator K, typ reflect.Type) {
	t.lookup[discriminator] = kindInfo[K]{
		discriminator: discriminator,
		typ:           typ,
	}
}

func (t KindMarshaler[K, D]) Marshal(val any) ([]byte, error) {
	if t.Formatted {
		return json.MarshalIndent(val, "", "    ")
	} else {
		return json.Marshal(val)
	}
}

func (t KindMarshaler[K, D]) Unmarshal(bs []byte) (any, error) {
	dv := reflect.New(t.discriminatorType)

	if err := json.Unmarshal(bs, dv.Interface()); err != nil {
		return nil, err
	}

	name := t.discriminatorFn(dv.Elem().Interface().(D))

	info, ok := t.lookup[name]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", name)
	}

	vv := reflect.New(info.typ)

	if err := json.Unmarshal(bs, vv.Interface()); err != nil {
		return nil, err
	}

	return vv.Elem().Interface(), nil
}
