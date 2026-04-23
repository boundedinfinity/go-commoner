package stringer

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func RemoveDiacritics[S ~string](s S) string {
	// Decompose Unicode characters (NFD)
	t := norm.NFD.String(string(s))

	// Remove all non-spacing marks (accents)
	var b strings.Builder
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
