package golang

import (
	"github.com/boundedinfinity/commons/runes"
	"github.com/boundedinfinity/commons/strings"
)

var (
	Keywords = []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct", "chan",
		"else", "goto", "package", "switch", "const",
		"fallthrough", "if", "range", "type", "continue",
		"for", "import", "return", "var", "any", "bool",
		"byte", "comparable", "complex64", "complex128",
		"error", "float32", "float64", "int", "int8",
		"int16", "int32", "int64", "rune", "string", "uint",
		"uint8", "uint16", "uint32", "uint64", "uintptr",
		"nil", "true", "false", "iota", "append", "cap",
		"close", "complex", "copy", "delete", "imag", "len",
		"make", "new", "panic", "print", "println", "real",
		"recover",
	}

	valid_first_chars    = []rune{}
	valid_nonfirst_chars = []rune{}
)

func init() {
	valid_first_chars = append(valid_first_chars, strings.UppercaseLetters...)
	valid_first_chars = append(valid_first_chars, strings.LowercaseLetters...)
	valid_first_chars = append(valid_first_chars, runes.UNDERSCORE)

	valid_nonfirst_chars = append(valid_nonfirst_chars, valid_first_chars...)
	valid_nonfirst_chars = append(valid_nonfirst_chars, strings.Numbers...)
}
