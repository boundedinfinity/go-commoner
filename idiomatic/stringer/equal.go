package stringer

import "strings"

func Equal[T ~string](s T, v string) bool {
	return string(s) == v
}

func EqualFold[T ~string](s T, v string) bool {
	return strings.EqualFold(string(s), v)
}

func EqualIgnoreCase[S, T ~string](s S, t T) bool {
	return strings.EqualFold(string(s), string(t))
}

func EqualIgnoreCaseTarget[S, T ~string](target S) func(T) bool {
	return func(b T) bool {
		return strings.EqualFold(string(target), string(b))
	}
}
