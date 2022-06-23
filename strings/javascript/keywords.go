package javascript

import "github.com/boundedinfinity/commons/slices"

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
)

func IsKeyword(s string) bool {
	return slices.Contains(Keywords, s)
}
