package langer

import "errors"

// https://go.dev/ref/spec#Identifiers

var Go = golang{}

type golang struct {
	keywords                             []string
	ErrIdentifierInvalidIsKeyword        error
	ErrIdentifierInvalidStartsWithNumber error
}

func (t golang) Identifier(s string) (string, error) {
	o := s

	for _, keyword := range t.keywords {
		if s == keyword {
			return o, t.ErrIdentifierInvalidIsKeyword
		}
	}

	return o, nil
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

	Go.ErrIdentifierInvalidIsKeyword = errors.New("invalid go identifier: is a keyword")
	Go.ErrIdentifierInvalidStartsWithNumber = errors.New("invalid go identifier: starts with number")
}
