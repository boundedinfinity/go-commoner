package stringer

import (
	"strings"
	"unicode"
)

func CompactSpace[S ~string](s S) string {
	first := IndexNonSpace(s)
	last := IndexNonSpaceRight(s)

	if first < 0 {
		first = 0
	}

	if last < 0 {
		last = len(s)
	}

	var b strings.Builder
	space := false

	for _, r := range s[:first] {
		b.WriteRune(r)
	}

	for _, r := range s[first:last] {
		if unicode.IsSpace(r) {
			if !space {
				b.WriteRune(r)
				space = true
			}
		} else {
			b.WriteRune(r)
			space = false
		}
	}

	for _, r := range s[last:] {
		b.WriteRune(r)
	}

	return b.String()
}

func FirstNonSpace[T ~string | ~byte](s T) int {
	for i, c := range string(s) {
		if !unicode.IsSpace(c) {
			return i
		}
	}

	return 0
}
