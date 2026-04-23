package stringer

import (
	"strings"
)

func Replace[T ~string](s T, new string, olds ...string) string {
	result := string(s)

	for _, old := range olds {
		result = strings.ReplaceAll(result, old, new)
	}

	return result
}

func ReplaceAll[S, O, N ~string](s S, old O, new N) string {
	return strings.ReplaceAll(string(s), string(old), string(new))
}
