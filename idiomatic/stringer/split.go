package stringer

import "strings"

func Split[T ~string](s T, sep string) []string {
	return strings.Split(string(s), sep)
}

func SplitAny[T ~string](s T, substrs ...string) []string {
	var results []string
	normal := string(s)
	var i int

	for _, substr := range substrs {
		if i = strings.Index(normal, substr); i > -1 {
			results = append(results, string(s[:i]))
			s = s[i+len(substr):]
		}
	}

	return results
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
