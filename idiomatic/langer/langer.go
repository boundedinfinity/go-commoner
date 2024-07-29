package langer

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

func new(name string) *langer {
	return &langer{
		name:                          name,
		ErrIdentifierInvalidIsKeyword: errorer.Errorf("invalid %s identifier: is a keyword", name),
	}
}

func (t *langer) WithCantStartWithNumber(value bool) *langer {
	t.canStartWithNumber = value
	return t
}

func (t *langer) WithKeywords(keywords ...string) *langer {
	t.keywords = append(t.keywords, keywords...)
	return t
}

func (t *langer) ResetKeywords(keywords ...string) *langer {
	t.keywords = keywords
	return t
}

func (t *langer) WithTranslations(translations map[string]string) *langer {
	mapper.MergeInto(t.translations, translations)
	return t
}

func (t *langer) WithTransformer(langFn ...func(s string) (string, error)) *langer {
	t.transformers = append(t.transformers, langFn...)
	return t
}

type langer struct {
	name                          string
	keywords                      []string
	translations                  map[string]string
	transformers                  []func(s string) (string, error)
	caserConversion               string
	canStartWithNumber            bool
	ErrIdentifierInvalidIsKeyword *errorer.Errorer
}

func (t langer) IsKeyword(s string) bool {
	return slicer.AnyOf(s, t.keywords...)
}

func (t *langer) WithCaserConversion(value string) *langer {
	t.caserConversion = value
	return t
}

func (t langer) MustIdentifier(s string) string {
	return t.MustIdentifierWithTranslation(s, map[string]string{})
}

func (t langer) MustIdentifierWithTranslation(s string, translate map[string]string) string {
	identifier, err := t.IdentifierWithTranslation(s, translate)

	if err != nil {
		panic(err)
	}

	return identifier
}

func (t langer) Identifier(s string) (string, error) {
	return t.IdentifierWithTranslation(s, map[string]string{})
}

func (t langer) IdentifierWithTranslation(s string, translate map[string]string) (string, error) {
	identifier := s
	var err error
	caserConversion := t.caserConversion

	identifier = utfer.RemoveNewlines(identifier)
	identifier = stringer.TrimSpace(identifier)

	if caserConversion == "" {
		caserConversion = "phrase-to-pascal"
	}

	for from, to := range translate {
		identifier = stringer.Replace(identifier, to, from)
	}

	for _, transformer := range t.transformers {
		if identifier, err = transformer(identifier); err != nil {
			return identifier, err
		}
	}

	identifier, err = caser.ConvertAs(identifier, caserConversion)
	if err != nil {
		return identifier, err
	}

	if t.IsKeyword(identifier) {
		err = t.ErrIdentifierInvalidIsKeyword.WithValue(identifier)
	}

	if !t.canStartWithNumber {
		if utfer.OneOf(identifier[0], utfer.Utf8.Numbers()) {
			identifier = "_" + identifier
		}
	}

	return identifier, err
}
