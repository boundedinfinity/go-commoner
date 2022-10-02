package stringer

import "strings"

func ToLower[T ~string](s T) string {
	return strings.ToLower(string(s))
}

func Lowercase[T ~string](s T) string {
	return ToLower(s)
}

func ToLowerFirst[T ~string](s T) string {
	if len(s) <= 1 {
		return ToLower(s)
	}

	f := string(s[0])
	r := string(s[1:])

	return ToLower(f) + r
}

func LowercaseFirst[T ~string](s T) string {
	return ToLowerFirst(s)
}
