package environmenter

import (
	"fmt"
	"os"
	"strings"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/stringer"
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

		if c.Patterns.Empty() {
			patterns = make([]string, 0)
		}

		patterns = c.Patterns.Get()
		patterns = append(patterns, ps...)
		c.Patterns = optioner.Some(patterns)
	}
}

type Config struct {
	Patterns optioner.Option[[]string]
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

		for _, format := range config.Patterns.Get() {
			pattern := fmt.Sprintf(format, key)
			o = strings.ReplaceAll(o, pattern, val)
		}
	}

	return o
}

func Sub(s string, args ...ConfigArg) string {
	var config Config

	for _, arg := range args {
		arg(&config)
	}

	return SubWithConfig(s, config)
}
