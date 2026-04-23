package stringer

import (
	"strings"
	"unicode"
)

func Index[S, P ~string](s S, substr P) int {
	return strings.Index(string(s), string(substr))
}

func IndexNonSpace[S ~string](s S) int {
	for i, r := range s {
		if !unicode.IsSpace(r) {
			return i
		}
	}
	return -1
}

func IndexNonSpaceRight[S ~string](s S) int {
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsSpace(rune(s[i])) {
			return i
		}
	}

	return -1
}
