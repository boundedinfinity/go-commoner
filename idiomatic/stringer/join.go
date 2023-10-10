package stringer

import gstring "strings"

func Join[T ~string](elems []T, sep string) string {
	ss := []string{}

	for _, elem := range elems {
		ss = append(ss, string(elem))
	}

	return gstring.Join(ss, sep)
}
