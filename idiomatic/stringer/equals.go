package stringer

import "strings"

func Equal[T ~string](s T, v string) bool {
	return string(s) == v
}

func EqualFold[T ~string](s T, v string) bool {
	return strings.EqualFold(string(s), v)
}

func EqualIgnoreCase[T ~string](s T, v string) bool {
	return EqualFold(s, v)
}
