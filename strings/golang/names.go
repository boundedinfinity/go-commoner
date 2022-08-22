package golang

// https://www.w3schools.com/go/go_variable_naming_rules.php
// https://go.dev/ref/spec#Keywords

import (
	gostrings "strings"

	"github.com/boundedinfinity/go-commoner/runes"
	"github.com/boundedinfinity/go-commoner/slices"
	"github.com/boundedinfinity/go-commoner/strings"
)

type NameConfig struct {
	ReplacementChar rune
	Public          bool
}

func Name(s string) string {
	return NameWithConfig(s, DefaultConfig())
}

func DefaultConfig() NameConfig {
	return NameConfig{
		ReplacementChar: runes.UNDERSCORE,
		Public:          true,
	}
}

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
		ns = strings.UppercaseFirst(ns)
	}

	return ns
}
