package golang

import "github.com/boundedinfinity/commons/slices"

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
)

func IsKeyword(s string) bool {
	return slices.Contains(Keywords, s)
}
