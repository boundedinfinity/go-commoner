package stringer

import "strings"

func ToUpper[T ~string](s T) string {
	return strings.ToUpper(string(s))
}

func ToUpperAll[T ~string](elems ...T) []string {
	results := make([]string, len(elems))

	for i := range elems {
		results[i] = strings.ToUpper(string(elems[i]))
	}

	return results
}

func ToUpperFirst[T ~string](s T) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(string(s))
	default:
		return strings.ToUpper(string(s[0])) + string(s[1:])
	}
}

func SliceToUpper[S ~string](s []S) []string {
	output := make([]string, len(s))

	for i, elem := range s {
		output[i] = strings.ToUpper(string(elem))
	}

	return output
}

func IsUpper[S ~string](s S) bool {
	return string(s) == strings.ToUpper(string(s))
}
