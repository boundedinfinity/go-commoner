package marshal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var (
	m = make(map[string]interface{})
)

func Register(v any) error {
	name := fqpp(v)

	if _, ok := m[name]; !ok {
		m[name] = v
	}

	return nil
}

func Unmarshal(data []byte) (any, error) {
	var w deserializationWrapper

	if err := json.Unmarshal(data, &w); err != nil {
		return nil, err
	}

	instance, ok := m[w.Name]

	if !ok {
		return nil, fmt.Errorf("unkown type %v", w.Name)
	}

	typ := reflect.TypeOf(instance)
	v := reflect.New(typ).Interface()

	if err := json.Unmarshal(w.Instance, v); err != nil {
		return nil, err
	}

	return v, nil
}

func UnmarshalTyped[T any](data []byte) (T, error) {
	iface, err := Unmarshal(data)

	if err != nil {
		var zero T
		return zero, err
	}

	instance, ok := iface.(*T)

	if !ok {
		var zero T
		return zero, fmt.Errorf("can't cast %v to %v", fqpp(iface), fqpp(zero))
	}

	return *instance, nil
}

type deserializationWrapper struct {
	Name     string          `json:"name,omitempty"`
	Instance json.RawMessage `json:"instance,omitempty"`
}

func Marshal(v any) ([]byte, error) {
	return json.Marshal(wrap(v))
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(wrap(v), prefix, indent)
}

type serializationWrapper struct {
	Name     string `json:"name,omitempty"`
	Instance any    `json:"instance,omitempty"`
}

func wrap(v any) serializationWrapper {
	return serializationWrapper{
		Name:     fqpp(v),
		Instance: v,
	}
}

func fqpp(t any) string {
	typ := reflect.TypeOf(t)
	return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
}
