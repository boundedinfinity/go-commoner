package stringer

import "strings"

func HasSuffx[T ~string, U ~string | ~byte](s T, prefixes ...U) bool {
	for _, prefix := range prefixes {
		if strings.HasSuffix(string(s), string(prefix)) {
			return true
		}
	}

	return false
}

func EndsWith[T ~string, U ~string | ~byte](s T, suffixes ...U) bool {
	return HasSuffx(s, suffixes...)
}
