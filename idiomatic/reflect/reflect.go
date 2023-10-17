package reflect

import (
	"fmt"
	"reflect"
)

func FullyQualifiedName(instance any) string {
	typ := reflect.TypeOf(instance)
	return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
}
