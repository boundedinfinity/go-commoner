package strings

import "strings"

func ContainsAny[T ~string](s T, vs ...string) bool {
	ns := string(s)

	for _, v := range vs {
		if strings.Contains(ns, v) {
			return true
		}
	}

	return false
}

func ContainsNone[T ~string](s T, vs ...string) bool {
	return !ContainsAny[T](s, vs...)
}

func ContainsAnyIgnoreCase[T ~string](s T, vs ...string) bool {
	ns := strings.ToLower(string(s))

	for _, v := range vs {
		if strings.Contains(ns, strings.ToLower(v)) {
			return true
		}
	}

	return false
}


func ContainsNoneIgnoreCase[T ~string](s T, vs ...string) bool {
	return !ContainsAnyIgnoreCase[T](s, vs...)
}