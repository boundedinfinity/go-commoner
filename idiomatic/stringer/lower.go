package stringer

import "strings"

func ToLower[T ~string](s T) string {
	return strings.ToLower(string(s))
}

func ToLowerAll[T ~string](elems ...T) []string {
	results := make([]string, len(elems))

	for i := range elems {
		results[i] = strings.ToLower(string(elems[i]))
	}

	return results
}

func ToLowerFirst[T ~string](s T) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToLower(string(s))
	default:
		return strings.ToLower(string(s[0])) + string(s[1:])
	}
}

func SliceToLower[S ~string](s []S) []string {
	output := make([]string, len(s))

	for i, elem := range s {
		output[i] = strings.ToLower(string(elem))
	}

	return output
}

func IsLower[S ~string](s S) bool {
	return string(s) == strings.ToLower(string(s))
}
