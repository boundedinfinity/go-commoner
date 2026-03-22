package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	ErrLanger = errorer.New("langer")
	errLanger = errorer.Func(ErrLanger)
)

func new(name string) *langer {
	return &langer{name: name}
}

func wrapStringWithErr(fn func(string) string) func(string) (string, error) {
	return func(s string) (string, error) {
		return fn(s), nil
	}
}

type langer struct {
	name                string
	reservedWords       []string
	replacements        map[string]string
	transformers        []func(s string) (string, error)
	startWithCharacters []string
}

func (t *langer) Copy(name string) *langer {
	return &langer{
		name:                name,
		reservedWords:       t.reservedWords[:],
		replacements:        mapper.Copy(t.replacements),
		transformers:        t.transformers[:],
		startWithCharacters: t.startWithCharacters[:],
	}
}

func (t *langer) StartsWith(characters ...string) *langer {
	t.startWithCharacters = characters
	return t
}

func (t *langer) ReservedWords(reservedWords ...string) *langer {
	t.reservedWords = reservedWords
	return t
}

func (t *langer) Replacements(replacements map[string]string) *langer {
	mapper.Merge(t.replacements, replacements)
	return t
}

func (t *langer) Transformers(fn ...func(s string) (string, error)) *langer {
	t.transformers = append(t.transformers, fn...)
	return t
}

func (t langer) IsReservedWord(s string) bool {
	return slicer.Contains(s, t.reservedWords...)
}

func (t langer) MustIdentifier(s string) string {
	identifier, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return identifier
}

func (t langer) Identifier(s string) (string, error) {
	identifier := s
	var err error

	for _, transformer := range t.transformers {
		if identifier, err = transformer(identifier); err != nil {
			return identifier, err
		}
	}

	if t.IsReservedWord(identifier) {
		err = errLanger("identifer %s (from '%s') is a reserved workd", identifier, s)
		return identifier, err
	}

	if len(t.startWithCharacters) > 0 && !stringer.StartsWith(identifier, t.startWithCharacters...) {
		identifier = "_" + identifier
	}

	return identifier, err
}
