package stringer

import "strings"

func Uppercase[T ~string](s T) string {
	return strings.ToUpper(string(s))
}

func UppercaseFirst[T ~string](s T) string {
	if len(s) <= 1 {
		return Uppercase(s)
	}

	f := string(s[0])
	r := string(s[1:])

	return Uppercase(f) + r
}
