package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func NewKind[K comparable, D any]() *KindMarshaler[K, D] {
	return &KindMarshaler[K, D]{
		lookup:   make(map[K]kindInfo[K]),
		handlers: make([]func(K, any), 0),
	}
}

type KindMarshaler[K comparable, D any] struct {
	discriminatorType reflect.Type
	discriminatorFn   func(D) K
	lookup            map[K]kindInfo[K]
	Formatted         bool
	handlers          []func(K, any)
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

func (t *KindMarshaler[K, D]) RegisterHandlerFn(handler func(K, any)) {
	t.handlers = append(t.handlers, handler)
}

func (t KindMarshaler[K, D]) Marshal(val any) ([]byte, error) {
	if t.Formatted {
		return json.MarshalIndent(val, "", "    ")
	} else {
		return json.Marshal(val)
	}
}

func (t KindMarshaler[K, D]) Unmarshal(bs []byte) error {
	dv := reflect.New(t.discriminatorType)

	if err := json.Unmarshal(bs, dv.Interface()); err != nil {
		return err
	}

	name := t.discriminatorFn(dv.Elem().Interface().(D))

	info, ok := t.lookup[name]

	if !ok {
		return fmt.Errorf("no type found for %v", name)
	}

	vv := reflect.New(info.typ)

	if err := json.Unmarshal(bs, vv.Interface()); err != nil {
		return err
	}

	for _, handler := range t.handlers {
		handler(name, vv.Elem().Interface())
	}

	return nil
}
