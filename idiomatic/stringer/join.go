package stringer

import "strings"

func Join[T any](sep string, elems ...T) string {
	strs := AsStrings(elems...)
	return strings.Join(strs, sep)
}
