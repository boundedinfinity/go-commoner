package stringer

import "strings"

func ToTitle[T ~string](s T) string {
	return strings.ToTitle(string(s))
}

func ToTitleAll[T ~string](elems ...T) []string {
	results := make([]string, len(elems))

	for i := range elems {
		results[i] = strings.ToTitle(string(elems[i]))
	}

	return results
}

func ToTitleFirst[T ~string](s T) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToTitle(string(s))
	default:
		return strings.ToTitle(string(s[0])) + string(s[1:])
	}
}

func SliceToTitle[S ~string](s []S) []string {
	output := make([]string, len(s))

	for i, elem := range s {
		output[i] = strings.ToTitle(string(elem))
	}

	return output
}

func IsTitle[S ~string](s S) bool {
	return string(s) == strings.ToTitle(string(s))
}
