package stringer

import "strings"

func Title[T ~string](s T) string {
	return strings.Title(string(s))
}

func Capitalize[T ~string](s T) string {
	return Title(s)
}
