package stringer

import "strings"

type Builder[T ~string] struct {
	strings.Builder
}

func (t *Builder[T]) WriteString(s T) (int, error) {
	return t.Builder.WriteString(string(s))
}
