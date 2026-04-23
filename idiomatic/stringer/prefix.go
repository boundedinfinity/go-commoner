package stringer

import "strings"

func HasPrefix[T ~string, U ~string](s T, prefixes ...U) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(string(s), string(prefix)) {
			return true
		}
	}

	return false
}

func StartsWith[T ~string, U ~string](s T, prefixes ...U) bool {
	return HasPrefix(s, prefixes...)
}
