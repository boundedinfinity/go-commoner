package marshaler

import (
	"encoding/json"
	"fmt"
	greflect "reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflect"
)

type unmarshalWrapper struct {
	Discriminator string          `json:"discriminator,omitempty"`
	Instance      json.RawMessage `json:"instance,omitempty"`
}

type marshalWrapper struct {
	Discriminator string `json:"discriminator,omitempty"`
	Instance      any    `json:"instance,omitempty"`
}

type MashalerManager struct {
	m map[string]greflect.Type
}

func New() *MashalerManager {
	return &MashalerManager{
		m: make(map[string]greflect.Type),
	}
}

func (t *MashalerManager) Register(instance any) error {
	name := reflect.FullyQualifiedName(instance)

	if _, ok := t.m[name]; !ok {
		t.m[name] = greflect.TypeOf(instance)
	}

	return nil
}

func (t MashalerManager) Marshal(instance any) ([]byte, error) {
	wrapped := marshalWrapper{
		Discriminator: reflect.FullyQualifiedName(instance),
		Instance:      instance,
	}

	return json.Marshal(wrapped)
}

func (t MashalerManager) MarshalIndent(instance any, prefix, indent string) ([]byte, error) {
	wrapped := marshalWrapper{
		Discriminator: reflect.FullyQualifiedName(instance),
		Instance:      instance,
	}

	return json.MarshalIndent(wrapped, prefix, indent)
}

func (t MashalerManager) Unmarshal(data []byte) (any, error) {
	var w unmarshalWrapper

	if err := json.Unmarshal(data, &w); err != nil {
		return nil, err
	}

	typ, ok := t.m[w.Discriminator]

	if !ok {
		return nil, fmt.Errorf("unkown type %v", w.Discriminator)
	}

	value := greflect.New(typ)
	instancePtr := value.Interface()

	if err := json.Unmarshal(w.Instance, instancePtr); err != nil {
		return nil, err
	}

	instance := greflect.Indirect(value).Interface()

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
