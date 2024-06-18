package reflecter

import (
	"fmt"
	"reflect"
)

// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7

func TypeQualifiedName[T any]() string {
	var t T
	return InstanceQualifiedName(t)
}

func InstanceQualifiedName(t any) string {
	typ := reflect.TypeOf(t)

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

func TypeBaseName[T any]() string {
	var t T
	return InstanceSimpleName(t)
}

func InstanceSimpleName(t any) string {
	typ := reflect.TypeOf(t)

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

func IsZero[T any](instance any) bool {
	var t T
	return reflect.DeepEqual(instance, t)
}
