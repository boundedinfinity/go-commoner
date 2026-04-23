package stringer

import "strings"

func HasSuffix[T ~string, U ~string | ~byte](s T, suffixes ...U) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(string(s), string(suffix)) {
			return true
		}
	}

	return false
}

func EndsWith[T ~string, U ~string | ~byte](s T, suffixes ...U) bool {
	return HasSuffix(s, suffixes...)
}
