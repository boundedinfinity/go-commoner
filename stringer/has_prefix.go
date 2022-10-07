package stringer

import "strings"

func HasPrefix[T ~string, U ~string](s T, prefix U) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func StartsWith[T ~string, U ~string](s T, prefix U) bool {
	return HasPrefix(s, prefix)
}
