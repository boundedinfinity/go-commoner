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
		WithKeywords("break", "default", "func", "interface", "select", "case", "defer",
			"go", "map", "struct", "chan", "else", "goto", "package", "switch",
			"const", "fallthrough", "if", "range", "type", "continue", "for",
			"import", "return", "var").
		WithLang(
			func(identifier string) (string, error) {
				identifier = utfer.RemoveNewlines(identifier)
				identifier = stringer.TrimSpace(identifier)
				identifier = utfer.RemoveSymbols(identifier)
				identifier = caser.PhraseToPascal(identifier)

				if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
					identifier = "_" + identifier
				}

				return identifier, nil
			},
		)
}
