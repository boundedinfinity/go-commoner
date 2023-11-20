package reflect

import (
	"fmt"
	"reflect"
)

var Instances = instances{}

type instances struct{}

func (t instances) QualifiedName(instance any) string {
	typ := reflect.TypeOf(instance)

	switch typ.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Array,
		reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.String, reflect.Struct:
		return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
	default:
		panic(fmt.Errorf("unsupported type: %v", typ.Kind()))
	}
}

func (t instances) SimpleName(instance any) string {
	typ := reflect.TypeOf(instance)

	switch typ.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Array,
		reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.String, reflect.Struct:
		return typ.Name()
	default:
		panic(fmt.Errorf("unsupported type: %v", typ.Kind()))
	}
}

func (t instances) IsZero(instance any) bool {
	return reflect.DeepEqual(
		instance,
		reflect.Zero(reflect.TypeOf(instance)).Interface(),
	)
}
