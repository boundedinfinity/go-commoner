package stringer

import "strings"

func HasSuffx[T ~string, U ~string](s T, prefix U) bool {
	return strings.HasSuffix(string(s), string(prefix))
}

func EndsWith[T ~string, U ~string](s T, prefix U) bool {
	return HasSuffx(s, prefix)
}
