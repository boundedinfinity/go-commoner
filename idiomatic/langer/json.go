package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

var Json *langer

func init() {
	javascript = new("json").
		WithLang(
			func(identifier string) (string, error) {
				identifier = utfer.RemoveNewlines(identifier)
				identifier = stringer.TrimSpace(identifier)
				identifier = caser.PhraseToKebabLower(identifier)

				if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
					identifier = "_" + identifier
				}

				return identifier, nil
			},
		)
}
