package stringer

import (
	"strings"
)

func Contains[T ~string, S ~string](s S, substr T) bool {
	return strings.Contains(string(s), string(substr))
}

func ContainsAny[T ~string, S ~string](s T, chars S) bool {
	return strings.ContainsAny(string(s), string(chars))
}

func ContainsIgnoreCase[T ~string, S ~string](s S, substr T) bool {
	return Contains(ToLower(s), ToLower(substr))
}

func ContainsAnySubString[T, S ~string](s S, substrs ...T) bool {
	native := string(s)

	for _, substr := range substrs {
		if strings.Contains(native, string(substr)) {
			return true
		}
	}

	return false
}

func ContainsAnyIgnoreCase[T ~string, S ~string](s S, substrs ...T) bool {
	return ContainsAnySubString(ToLower(s), TransformAllFunc(substrs, ToLower)...)
}

func ContainsNone[T ~string, S ~string](s S, substrs ...T) bool {
	return !ContainsAnySubString(s, substrs...)
}

func ContainsNoneIgnoreCase[T ~string, S ~string](s S, substrs ...T) bool {
	return !ContainsAnyIgnoreCase(s, substrs...)
}
