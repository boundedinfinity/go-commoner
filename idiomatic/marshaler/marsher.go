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

type Marshaler struct {
	types map[string]marshalerContext
}

func (t *Marshaler) Register(item any) {
	typ := reflect.TypeOf(item)

	ctx := marshalerContext{
		typ: typ,
	}

	t.types[typ.Name()] = ctx
}

func (t Marshaler) Marshal(item any) ([]byte, error) {
	name := reflect.TypeOf(item).Name()

	if _, ok := t.types[name]; !ok {
		return nil, fmt.Errorf("no type found for %v", name)
	}

	wrapped := marshalWrapper{
		Type:  name,
		Value: item,
	}

	return json.Marshal(wrapped)
}

func (t Marshaler) Unmarshal(data []byte) (any, error) {
	var err error
	var wrapped unmarshalWrapper

	if err = json.Unmarshal(data, &wrapped); err != nil {
		return nil, err
	}

	ctx, ok := t.types[wrapped.Type]

	if !ok {
		return nil, fmt.Errorf("no type found for %v", wrapped.Type)
	}

	rf := reflect.New(ctx.typ)

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
