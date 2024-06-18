package langer

import "github.com/boundedinfinity/go-commoner/idiomatic/utfer"

var Sql = _sql{}

type _sql struct{}

func (t _sql) MustIdentifier(s string) string {
	identifier, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return identifier
}

func (t _sql) Identifier(s string) (string, error) {
	identifier := s

	if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
		identifier = "_" + identifier
	}

	identifier = utfer.ReplaceNewlines(identifier, byte(utfer.UNDERSCORE))
	identifier = utfer.ReplaceSpaces(identifier, byte(utfer.UNDERSCORE))
	identifier = utfer.ReplaceSymbols(identifier, byte(utfer.UNDERSCORE))

	return identifier, nil
}
