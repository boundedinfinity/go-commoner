package marshaler

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MashalerManager struct {
	m map[string]reflect.Type
}

func New() *MashalerManager {
	return &MashalerManager{
		m: make(map[string]reflect.Type),
	}
}

func (t *MashalerManager) Register(instance any) error {
	name := fqpp(instance)

	if _, ok := t.m[name]; !ok {
		t.m[name] = reflect.TypeOf(instance)
	}

	return nil
}

func (t MashalerManager) Unmarshal(data []byte) (any, error) {
	var w deserializationWrapper

	if err := json.Unmarshal(data, &w); err != nil {
		return nil, err
	}

	typ, ok := t.m[w.Name]

	if !ok {
		return nil, fmt.Errorf("unkown type %v", w.Name)
	}

	// instance := reflect.New(typ).Elem().Interface()
	// instance := reflect.Indirect(reflect.New(typ)).Elem().Interface()
	// instance := reflect.Indirect(reflect.New(typ)).Interface()
	// instance := reflect.make
	// instance := reflect.Zero(typ).Elem().Interface()
	// instance := reflect.Zero(typ).Interface()

	value := reflect.New(typ)
	instancePtr := value.Interface()

	if err := json.Unmarshal(w.Instance, instancePtr); err != nil {
		return nil, err
	}

	instance := reflect.Indirect(value).Interface()

	return instance, nil
}

// func UnmarshalTyped[T any](data []byte) (T, error) {
// 	iface, err := Unmarshal(data)

// 	if err != nil {
// 		var zero T
// 		return zero, err
// 	}

// 	instance, ok := iface.(*T)

// 	if !ok {
// 		var zero T
// 		return zero, fmt.Errorf("can't cast %v to %v", fqpp(iface), fqpp(zero))
// 	}

// 	return *instance, nil
// }

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

func fqpp(instance any) string {
	typ := reflect.TypeOf(instance)
	return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
}
