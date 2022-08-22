package javascript

import (
	"github.com/boundedinfinity/go-commoner/runes"
	"github.com/boundedinfinity/go-commoner/strings"
)

var (
	Keywords = []string{
		"break", "case", "catch", "NaN", "undefined",
		"continue", "debugger", "default", "delete",
		"do", "else", "finally", "for", "function",
		"if", "in", "instanceof", "new", "return",
		"switch", "this", "throw", "try", "typeof",
		"var", "void", "while", "with", "class",
		"const", "enum", "export", "extends",
		"import", "super", "implements", "interface",
		"let", "package", "private", "protected",
		"public", "static", "yield", "null", "true",
		"false", "int", "byte", "char", "goto",
		"final", "float", "short", "long", "double",
		"native", "throws", "boolean", "abstract",
		"volatile", "transient", "synchronized", "var",
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
