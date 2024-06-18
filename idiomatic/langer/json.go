package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

var Json = _json{}

type _json struct{}

func (t _json) MustIdentifier(s string) string {
	identifier, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return identifier
}

func (t _json) Identifier(s string) (string, error) {
	identifier := s

	identifier = utfer.RemoveNewlines(identifier)
	identifier = caser.PhraseToKebabLower(identifier)

	if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
		identifier = "_" + identifier
	}

	return identifier, nil
}
