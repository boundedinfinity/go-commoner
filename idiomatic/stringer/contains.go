package stringer

import "strings"

func ContainsAny[T ~string, S ~string](s S, items ...T) bool {
	ns := string(s)

	for _, item := range items {
		if strings.Contains(ns, string(item)) {
			return true
		}
	}

	return false
}

func ContainsNone[T ~string, S ~string](s S, items ...T) bool {
	return !ContainsAny(s, items...)
}

func ContainsAnyIgnoreCase[T ~string, S ~string](s S, items ...T) bool {
	ns := strings.ToLower(string(s))

	for _, v := range items {
		if strings.Contains(ns, strings.ToLower(string(v))) {
			return true
		}
	}

	return false
}

func ContainsNoneIgnoreCase[T ~string, S ~string](s S, items ...T) bool {
	return !ContainsAnyIgnoreCase(s, items...)
}
