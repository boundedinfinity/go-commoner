package environmenter

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/sheller"
)

func New(mappings ...map[string]string) *Environment {
	return Shell(sheller.Shells.Unknown, mappings...)
}

func Shell(shell sheller.Shell, mappings ...map[string]string) *Environment {
	this := &Environment{
		shell:    shell,
		mapping:  mapper.Append(mappings...),
		reversed: make(map[string]string),
	}

	for key, value := range this.mapping {
		var patterns []string

		switch this.shell {
		case sheller.Shells.Unknown:
			patterns = createDetectedPatterns(key)
		default:
			patterns = createShellPatterns(this.shell, key)
		}

		for _, pattern := range patterns {
			this.reversed[pattern] = value
		}
	}

	return this
}

type Environment struct {
	shell    sheller.Shell
	mapping  map[string]string
	reversed map[string]string
}

func (this Environment) SubstitueAll(items ...string) []string {
	subs := make([]string, len(items))

	for i := range items {
		subs[i] = this.Substitue(items[i])
	}

	return subs
}

func (this Environment) Substitue(item string) string {
	var ok bool
	item, ok = this.substitueOk(item)

	for ok {
		item, ok = this.substitueOk(item)
	}

	return item
}

func (this Environment) substitueOk(item string) (string, bool) {
	var ok bool

	for key, value := range this.reversed {
		if strings.Contains(item, key) {
			item = strings.ReplaceAll(item, key, value)
			ok = true
		}
	}

	return item, ok
}

func createDetectedPatterns(name string) []string {
	return createShellPatterns(sheller.Detect(), name)
}

func createShellPatterns(shell sheller.Shell, name string) []string {
	var patterns []string

	switch shell {
	case sheller.Shells.Bash:
		patterns = append(patterns,
			`$`+name,
			`${`+name+`}`,
		)
	case sheller.Shells.Csh, sheller.Shells.Fish, sheller.Shells.Sh, sheller.Shells.Zsh:
		patterns = append(patterns,
			`$`+name,
		)
	case sheller.Shells.Powershell:
		patterns = append(patterns,
			`$env:`+name,
		)
	case sheller.Shells.Cmd:
		patterns = append(patterns,
			`%%`+name+`%%`,
		)
	}

	return patterns
}
