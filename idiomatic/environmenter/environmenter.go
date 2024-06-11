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

func (t *Environmenter) subWithMap(elem string, envMap map[string]string) string {
	replaced := elem

	for key, val := range envMap {
		for _, format := range t.patterns {
			pattern := fmt.Sprintf(format, key)
			replaced = strings.ReplaceAll(replaced, pattern, val)
		}
	}

	return replaced
}

func (t *Environmenter) Sub(elem string) string {
	return t.subWithMap(elem, t.parseEnv())
}

func (t *Environmenter) SubAll(elems ...string) []string {
	var replaced []string
	envMap := t.parseEnv()

	for _, elem := range elems {
		replaced = append(replaced, t.subWithMap(elem, envMap))
	}

	return replaced
}

func Sub(elem string) string {
	return New().Sub(elem)
}

func SubAll(elems ...string) []string {
	return New().SubAll(elems...)
}
