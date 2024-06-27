package environmenter

import (
	"fmt"
	"os"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	default_patterns = []string{
		"$%v",
		"${%v}",
		"${env:%v}",
	}

	default_envFileDirs = []string{
		"$HOME",
		"$HOME/.config",
		".",
	}

	default_envFileNames = []string{
		".env",
		"env",
	}
)

type EnvironmenterContext struct {
	Value  string
	Source string
}

func New() *Environmenter {
	t := &Environmenter{
		includeEnvironment: true,
		includeEnvFile:     true,
		environment:        map[string]EnvironmenterContext{},
		overrides:          map[string]EnvironmenterContext{},
	}

	t.
		WithPatterns(default_patterns...).
		WithSep("=").
		WithEnvFileDirs(default_envFileDirs...).
		WithEnvFileNames(default_envFileNames...)

	return t
}

type Environmenter struct {
	environment        map[string]EnvironmenterContext
	patterns           []string
	overrides          map[string]EnvironmenterContext
	sep                string
	includeEnvironment bool
	includeEnvFile     bool
	envFileDirs        []string
	envFileNames       []string
}

func (t *Environmenter) Process() error {
	if t.includeEnvironment {
		mapper.MergeInto(t.environment, t.parseLines("environment", os.Environ()...))
	}

	if t.includeEnvironment {
		for _, path := range t.generateEnvPaths() {
			if pather.Files.Exists(path) {
				data, err := os.ReadFile(path)

				if err != nil {
					return err
				}

				mapper.MergeInto(t.environment, t.parseData(path, string(data)))
			}
		}
	}

	mapper.MergeInto(t.environment, t.overrides)

	return nil
}

func (t Environmenter) generateEnvPaths() []string {
	var paths []string

	for _, dir := range t.envFileDirs {
		for _, name := range t.envFileNames {
			paths = append(paths, pather.Paths.Join(dir, name))
		}
	}

	return paths
}

func (t *Environmenter) parseLines(source string, lines ...string) map[string]EnvironmenterContext {
	found := map[string]EnvironmenterContext{}

	for _, line := range lines {
		comps := stringer.Split(line, t.sep)

		if len(comps) != 2 {
			continue
		}

		key := comps[0]
		val := comps[1]
		found[key] = EnvironmenterContext{
			Value:  val,
			Source: source,
		}
	}

	return found
}

func (t *Environmenter) parseData(source string, data string) map[string]EnvironmenterContext {
	return t.parseLines(source, stringer.Split(data, "\n")...)
}

func (t *Environmenter) convertMap(source string, inputs map[string]string) map[string]EnvironmenterContext {
	m := map[string]EnvironmenterContext{}

	for k, v := range inputs {
		m[k] = EnvironmenterContext{
			Value:  v,
			Source: source,
		}
	}

	return m
}

// WithSep overrides the separater character split key value pairs of the following form:
//
// KEY=VALUE
func (t *Environmenter) WithSep(splitChar string) *Environmenter {
	t.sep = splitChar
	return t
}

// WithData overrides environment variable key value pairs of the form:
//
// KEY=VALUE
func (t *Environmenter) WithData(source, data string) *Environmenter {
	mapper.MergeInto(t.overrides, t.parseData(source, data))
	return t
}

// WithLines overrides environment variable key value pairs of the form:
//
// KEY=VALUE
func (t *Environmenter) WithLines(source string, lines ...string) *Environmenter {
	mapper.MergeInto(t.overrides, t.parseLines(source, lines...))
	return t
}

// Withdata overrides environment variable key value pairs of the form:
//
// KEY=VALUE
func (t *Environmenter) WithMap(source string, environment map[string]string) *Environmenter {
	mapper.MergeInto(t.overrides, t.convertMap(source, environment))
	return t
}

// WithPatterns updates the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) WithPatterns(patterns ...string) *Environmenter {
	t.patterns = append(t.patterns, patterns...)
	return t
}

// OverridePatterns overrides the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) OverridePatterns(patterns ...string) *Environmenter {
	t.patterns = patterns
	return t
}

// WithPatterns updates the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) WithEnvFileDirs(paths ...string) *Environmenter {
	t.envFileDirs = append(t.envFileDirs, paths...)
	return t
}

// OverrideEnvFilePaths updates the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) OverrideEnvFilePaths(paths ...string) *Environmenter {
	t.envFileDirs = paths
	return t
}

// WithPatterns updates the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) WithEnvFileNames(names ...string) *Environmenter {
	t.envFileNames = append(t.envFileNames, names...)
	return t
}

// OverrideEnvFilePaths updates the list of patterns used to extract the environment
// varialbe names
func (t *Environmenter) OverrideEnvFileNames(names ...string) *Environmenter {
	t.envFileNames = names
	return t
}

func (t *Environmenter) Substitue(elem string) string {
	replaced, _ := t.SubstitueOk(elem)
	return replaced
}

func (t *Environmenter) SubstitueAll(elems ...string) []string {
	var replaced []string

	for _, elem := range elems {
		replaced = append(replaced, t.Substitue(elem))
	}

	return replaced
}

func (t *Environmenter) SubstitueOk(elem string) (string, bool) {
	replaced := elem

	for key, ctx := range t.environment {
		for _, format := range t.patterns {
			pattern := fmt.Sprintf(format, key)
			replaced = strings.ReplaceAll(replaced, pattern, ctx.Value)
		}
	}

	return replaced, replaced != elem
}

// Get returns the value of the environment variable for the given name.  If no value is found,
// returns an empty string.
func (t *Environmenter) Get(name string) (EnvironmenterContext, bool) {
	val, ok := t.environment[name]
	return val, ok
}

// Get returns a map of the values of the environment variable names.  If no value is found,
// returns an empty string.
func (t *Environmenter) GetAll(names ...string) map[string]EnvironmenterContext {
	environment := map[string]EnvironmenterContext{}

	for _, name := range names {
		if name == "" {
			continue
		}

		if found, ok := t.Get(name); ok {
			environment[name] = found
		}
	}

	return environment
}

func SubstitueOk(elem string) (string, bool) {
	return New().SubstitueOk(elem)
}

func Substitue(elem string) string {
	return New().Substitue(elem)
}

func SubAll(elems ...string) []string {
	return New().SubstitueAll(elems...)
}
