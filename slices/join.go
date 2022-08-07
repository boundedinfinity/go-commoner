package slices

import (
	"fmt"
	gstring "strings"
)

func Join[T any](elems []T, sep string) string {
	fn := func(el T) string {
		return fmt.Sprintf("%v", el)
	}

	ss := Map(elems, fn)
	return gstring.Join(ss, sep)
}
