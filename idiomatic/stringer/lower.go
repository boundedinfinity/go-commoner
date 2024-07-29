package stringer

import "strings"

func Lowercase[T ~string](s T) string {
	return strings.ToLower(string(s))
}

func LowercaseAll[T ~string](elems ...T) []string {
	var results []string

	for _, elem := range elems {
		results = append(results, Lowercase(elem))
	}

	return results
}

func LowercaseFirst[T ~string](s T) string {
	if len(s) <= 1 {
		return Lowercase(s)
	}

	f := string(s[0])
	r := string(s[1:])

	return Lowercase(f) + r
}
