package reflect

import (
	"fmt"
	"reflect"
)

func FullyQualifiedName(instance any) string {
	typ := reflect.TypeOf(instance)

	if typ == nil {
		return "nil/nil"
	} else {
		return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
	}
}
