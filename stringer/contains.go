package stringer

import "strings"

func ContainsAny[T ~string, S ~string](vs []T, s S) bool {
	ns := string(s)

	for _, v := range vs {
		if strings.Contains(ns, string(v)) {
			return true
		}
	}

	return false
}

func ContainsNone[T ~string, S ~string](vs []T, s S) bool {
	return !ContainsAny(vs, s)
}

func ContainsAnyIgnoreCase[T ~string, S ~string](vs []T, s S) bool {
	ns := strings.ToLower(string(s))

	for _, v := range vs {
		if strings.Contains(ns, strings.ToLower(string(v))) {
			return true
		}
	}

	return false
}

func ContainsNoneIgnoreCase[T ~string, S ~string](vs []T, s S) bool {
	return !ContainsAnyIgnoreCase(vs, s)
}
