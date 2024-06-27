package stringer

import (
	"strings"
)

func TabToSpace[T ~string | ~byte](s T, tabWidth int) string {
	var builder strings.Builder

	for _, c := range string(s) {
		if c == '\t' {
			builder.WriteString(strings.Repeat(" ", tabWidth))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

func SpaceToTab[T ~string | ~byte](s T, tabWidth int) string {
	return string(s)
}
