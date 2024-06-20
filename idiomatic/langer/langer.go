package langer

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func new(name string) *langer {
	return &langer{
		name:                          name,
		ErrIdentifierInvalidIsKeyword: errorer.Errorf("invalid %s identifier: is a keyword", name),
	}
}

func (t *langer) WithKeywords(keywords ...string) *langer {
	t.keywords = append(t.keywords, keywords...)
	return t
}

func (t *langer) WithTranslations(translations map[string]string) *langer {
	mapper.MergeInto(t.translations, translations)
	return t
}

func (t *langer) WithLang(langFn ...func(s string) (string, error)) *langer {
	t.langFns = append(t.langFns, langFn...)
	return t
}

type langer struct {
	name                          string
	keywords                      []string
	translations                  map[string]string
	langFns                       []func(s string) (string, error)
	ErrIdentifierInvalidIsKeyword *errorer.Errorer
}

func (t langer) IsKeyword(s string) bool {
	return slicer.Contains(s, t.keywords...)
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

	for from, to := range translate {
		identifier = stringer.Replace(identifier, to, from)
	}

	for _, langFn := range t.langFns {
		if identifier, err = langFn(identifier); err != nil {
			return identifier, err
		}
	}

	if t.IsKeyword(identifier) {
		err = t.ErrIdentifierInvalidIsKeyword.WithValue(identifier)
	}

	return identifier, err
}
