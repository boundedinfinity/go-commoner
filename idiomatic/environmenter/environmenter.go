package environmenter

import (
	"fmt"
	"os"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	default_patterns = []string{
		"$%v",
		"${%v}",
		"${env:%v}",
	}
)

type ConfigArg func(*Config)

func Pattern(p string) ConfigArg {
	return Patterns([]string{p})
}

func Patterns(ps []string) ConfigArg {
	return func(c *Config) {
		var patterns []string

		if c.Patterns == nil {
			patterns = make([]string, 0)
		}

		patterns = append(patterns, c.Patterns...)
		patterns = append(patterns, ps...)
		c.Patterns = patterns
	}
}

type Config struct {
	Patterns []string
}

func SubWithConfig(s string, config Config) string {
	o := s

	for _, env := range os.Environ() {
		comps := stringer.Split(env, "=")

		if len(comps) != 2 {
			continue
		}

		key := comps[0]
		val := comps[1]

		for _, format := range config.Patterns {
			pattern := fmt.Sprintf(format, key)
			o = strings.ReplaceAll(o, pattern, val)
		}
	}

	return o
}

func SubWithArgs(s string, args ...ConfigArg) string {
	var config Config

	if len(args) == 0 {
		for _, p := range default_patterns {
			args = append(args, Pattern(p))
		}
	}

	for _, arg := range args {
		arg(&config)
	}

	return SubWithConfig(s, config)
}

func Sub(s string) string {
	args := make([]ConfigArg, 0)

	for _, p := range default_patterns {
		args = append(args, Pattern(p))
	}

	return SubWithArgs(s, args...)
}
