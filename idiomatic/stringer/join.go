package stringer

import "strings"

func Join[T ~string](sep string, items ...T) string {
	ss := []string{}

	for _, item := range items {
		ss = append(ss, string(item))
	}

	return strings.Join(ss, sep)
}
