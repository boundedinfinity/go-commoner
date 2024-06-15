package langer

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

// https://mathiasbynens.be/notes/javascript-identifiers

var JavaScript = javaScript{}

type javaScript struct {
	keywords                             []string
	ErrIdentifierInvalidIsKeyword        *errorer.Errorer
	ErrIdentifierInvalidStartsWithNumber *errorer.Errorer
}

func (t javaScript) IsKeyword(s string) bool {
	return slicer.Contains(s, t.keywords...)
}

func (t javaScript) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t javaScript) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}

func init() {
	JavaScript.keywords = []string{
		"break", "case", "catch", "continue", "debugger", "default", "delete", "do", "else", "finally", "for",
		"function", "if", "in", "instanceof", "new", "return", "switch", "this", "throw", "try", "typeof", "var",
		"void", "while", "with", "class", "const", "enum", "export", "extends", "import", "super", "implements",
		"interface", "let", "package", "private", "protected", "public", "static", "yield",
	}

	JavaScript.ErrIdentifierInvalidIsKeyword = errorer.New("invalid javascript identifier: is a keyword")
	JavaScript.ErrIdentifierInvalidStartsWithNumber = errorer.New("invalid javascript identifier: starts with number")
}
