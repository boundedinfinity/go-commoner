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

func New() *Environmenter {
	t := &Environmenter{}
	return t.
		WithPatterns(default_patterns...).
		WithSplitChar("=")
}

type Environmenter struct {
	patterns    []string
	splitChar   string
	environment []string
}

func (t *Environmenter) WithSplitChar(splitChar string) *Environmenter {
	t.splitChar = splitChar
	return t
}

func (t *Environmenter) WithPatterns(patterns ...string) *Environmenter {
	t.patterns = append(t.patterns, patterns...)
	return t
}

func (t *Environmenter) WithEnvironment(environment ...string) *Environmenter {
	t.environment = append(t.environment, environment...)
	return t
}

func (t *Environmenter) ResetPatterns(patterns ...string) *Environmenter {
	t.patterns = patterns
	return t
}

func (t *Environmenter) parseEnv() map[string]string {
	env := append(os.Environ(), t.environment...)
	envMap := map[string]string{}

	for _, env := range env {
		comps := stringer.Split(env, t.splitChar)

		if len(comps) != 2 {
			continue
		}

		key := comps[0]
		val := comps[1]
		envMap[key] = val
	}

	return envMap
}

func (t *Environmenter) subWithMap(item string, envMap map[string]string) string {
	replaced := item

	for key, val := range envMap {
		for _, format := range t.patterns {
			pattern := fmt.Sprintf(format, key)
			replaced = strings.ReplaceAll(replaced, pattern, val)
		}
	}

	return replaced
}

func (t *Environmenter) Sub(item string) string {
	return t.subWithMap(item, t.parseEnv())
}

func (t *Environmenter) SubAll(items ...string) []string {
	var replaced []string
	envMap := t.parseEnv()

	for _, item := range items {
		replaced = append(replaced, t.subWithMap(item, envMap))
	}

	return replaced
}

func Sub(item string) string {
	return New().Sub(item)
}

func SubAll(items ...string) []string {
	return New().SubAll(items...)
}
