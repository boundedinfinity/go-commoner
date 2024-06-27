package stringer

import "unicode"

func FirstNonSpace[T ~string | ~byte](s T) int {
	for i, c := range string(s) {
		if !unicode.IsSpace(c) {
			return i
		}
	}

	return 0
}
