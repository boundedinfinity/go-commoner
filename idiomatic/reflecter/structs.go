package reflecter

import (
	"errors"
	"reflect"
)

var (
	ErrStructFieldNamesNotStruct = errors.New("no struct")
)

func StructFieldNames[T any]() ([]string, error) {
	var t T
	var names []string
	typeInfo := reflect.TypeOf(t)

	switch typeInfo.Kind() {
	case reflect.Struct:
		for i := 0; i < typeInfo.NumField(); i++ {
			names = append(names, typeInfo.Field(i).Name)
		}
	default:
		return names, ErrStructFieldNamesNotStruct
	}

	return names, nil
}

func StructFieldNamesAndTypes[T any]() (map[string]string, error) {
	info := map[string]string{}
	var t T
	typeInfo := reflect.TypeOf(t)

	switch typeInfo.Kind() {
	case reflect.Struct:
		for i := 0; i < typeInfo.NumField(); i++ {
			name := typeInfo.Field(i).Name
			typ := typeInfo.Field(i).Type.String()

			info[name] = typ
		}
	default:
		return info, ErrStructFieldNamesNotStruct
	}

	return info, nil
}
