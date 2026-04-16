package stringer

import "strings"

func ToCapital[T ~string](s T) string {
	if len(s) == 0 {
		return ""
	}

	words := strings.Split(string(s), " ")

	for i := range words {
		words[i] = ToUpperFirst(words[i])
	}

	return strings.Join(words, " ")
}

func ToCapitalAll[T ~string](elems ...T) []string {
	results := make([]string, len(elems))

	for i := range elems {
		results[i] = ToCapital(elems[i])
	}

	return results
}
