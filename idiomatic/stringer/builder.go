package stringer

import (
	"fmt"
	"strings"
)

type Builder[T ~string] struct {
	strings.Builder
}

func (t *Builder[T]) WriteString(s T) (int, error) {
	return t.Builder.WriteString(string(s))
}

func (t *Builder[T]) WriteStringf(format string, a ...any) (int, error) {
	return t.Builder.WriteString(fmt.Sprintf(format, a...))
}
