package stringer

import (
	"fmt"
	"strings"
)

type Builder struct {
	builder strings.Builder
}

func (this *Builder) WriteStringf(format string, a ...any) {
	this.builder.WriteString(fmt.Sprintf(format, a...))
}

func (this *Builder) WriteString(s string) {
	this.builder.WriteString(s)
}

func (this *Builder) WriteRune(r rune) {
	this.builder.WriteRune(r)
}

func (this Builder) String() string {
	return this.builder.String()
}
