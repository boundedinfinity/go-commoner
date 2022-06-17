package strings

import gostrings "strings"

func ToUpper[T ~string](s T) string {
	return gostrings.ToUpper(string(s))
}

func Uppercase[T ~string](s T) string {
	return ToUpper(s)
}

func ToUpperFirst[T ~string](s T) string {
	if len(s) <= 1 {
		return ToUpper(s)
	}

	f := string(s[0])
	r := string(s[1:])

	return ToUpper(f) + r
}

func UppercaseFirst[T ~string](s T) string {
	return ToUpperFirst(s)
}
