package langer

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

// https://go.dev/ref/spec#Identifiers

var Go = golang{}

type golang struct {
	keywords                             []string
	ErrIdentifierInvalidIsKeyword        *errorer.Errorer
	ErrIdentifierInvalidStartsWithNumber *errorer.Errorer
}

func (t golang) IsKeyword(s string) bool {
	return slicer.Contains(s, t.keywords...)
}

func (t golang) MustIdentifier(s string) string {
	identifier, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return identifier
}

func (t golang) Identifier(s string) (string, error) {
	identifier := s

	identifier = utfer.RemoveNewlines(identifier)
	identifier = utfer.RemoveSymbols(identifier)
	identifier = caser.PhraseToPascal(identifier)

	if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
		identifier = "_" + identifier
	}

	if t.IsKeyword(identifier) {
		return identifier, t.ErrIdentifierInvalidIsKeyword.WithValue(identifier)
	}

	return identifier, nil
}

func init() {
	Go.keywords = []string{
		"break", "default", "func", "interface", "select", "case", "defer",
		"go", "map", "struct", "chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type", "continue", "for",
		"import", "return", "var",
	}

	Go.ErrIdentifierInvalidIsKeyword = errorer.New("invalid go identifier: is a keyword")
	Go.ErrIdentifierInvalidStartsWithNumber = errorer.New("invalid go identifier: starts with number")
}
