package stringer

import "strings"

func Join[T ~string](sep string, elems ...T) string {
	ss := []string{}

	for _, elem := range elems {
		ss = append(ss, string(elem))
	}

	return strings.Join(ss, sep)
}
