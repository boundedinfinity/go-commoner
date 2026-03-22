package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

// https://go.dev/ref/spec#Identifiers

var Go *langer

func init() {
	Go = new("go").
		ReservedWords(
			"break", "default", "func", "interface", "select", "case", "defer",
			"go", "map", "struct", "chan", "else", "goto", "package", "switch",
			"const", "fallthrough", "if", "range", "type", "continue", "for",
			"import", "return", "var",
		).StartsWith(
		"_", "a",
	).Transformers(
		wrapStringWithErr(utfer.RemoveNewlines),
		wrapStringWithErr(stringer.TrimSpace),
		wrapStringWithErr(caser.PhraseToPascal),
	)
}
