package javascript

// https://mathiasbynens.be/notes/javascript-identifiers-es6

import (
	gostrings "strings"

	"github.com/boundedinfinity/go-commoner/runes"
	"github.com/boundedinfinity/go-commoner/slices"
)

type NameConfig struct {
	ReplacementChar rune
	SeparatorChar   string
	Public          bool
}

func Name(s string) string {
	return NameWithConfig(s, DefaultConfig())
}

func DefaultConfig() NameConfig {
	return NameConfig{
		ReplacementChar: runes.UNDERSCORE,
		Public:          true,
		SeparatorChar:   ".",
	}
}

// func JoinWithConfig(config NameConfig, ss ...string) string {
// 	var comps1 []string

// 	:= strings.Split()

// 	return strings.Join(comps1, "")
// }

func NameWithConfig(s string, config NameConfig) string {
	if len(s) < 1 {
		return s
	}

	var builder gostrings.Builder

	for i, r := range s {
		if i == 0 {
			if !slices.Contains(valid_first_chars, r) {
				builder.WriteRune(config.ReplacementChar)
			}

			builder.WriteRune(r)
		} else {
			if slices.Contains(valid_nonfirst_chars, r) {
				builder.WriteRune(r)
			} else {
				builder.WriteRune(config.ReplacementChar)
			}
		}
	}

	ns := builder.String()

	if IsKeyword(ns) {
		ns = string(config.ReplacementChar) + ns
	}

	if config.Public {
		// ns = strings.UppercaseFirst(ns)
	}

	return ns
}
