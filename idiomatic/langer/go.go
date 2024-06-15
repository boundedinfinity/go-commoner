package langer

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
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

func (t golang) Identifier(s string) (string, error) {
	identifier := s

	if t.IsKeyword(s) {
		return identifier, t.ErrIdentifierInvalidIsKeyword.WithValue(identifier)
	}

	// if stringer.HasPrefix(identifier, utf.Utf7.Numbers()...) {
	// 	return identifier, t.ErrIdentifierInvalidStartsWithNumber.WithValue(identifier)
	// }

	identifier = stringer.ReplaceInList(s, []string{" ", "-"}, "_")

	// o := s

	// for _, keyword := range t.keywords {
	// 	if s == keyword {
	// 		return o, t.ErrIdentifierInvalidIsKeyword
	// 	}
	// }

	return identifier, nil
}

func (t golang) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
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
