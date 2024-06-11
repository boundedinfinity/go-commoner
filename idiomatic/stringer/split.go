package stringer

import "strings"

func Split[T ~string](s T, sep string) []string {
	return strings.Split(string(s), sep)
}

func SplitAny[T ~string](s T, elems ...string) []string {
	return SplitFn(s, func(x string) bool {
		for _, elem := range elems {
			if elem == string(x) {
				return true
			}
		}

		return false
	})
}

func SplitFn[T ~string](s T, fn func(string) bool) []string {
	var o []string
	start := 0

	for i, x := range s {
		if fn(string(x)) {
			o = append(o, string(s[start:i]))
			start = i + 1
		}
	}

	o = append(o, string(s[start:]))

	return o
}
