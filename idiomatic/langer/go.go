package langer

// https://go.dev/ref/spec#Identifiers

var Go *langer

func init() {
	Go = new("go").
		WithKeywords("break", "default", "func", "interface", "select", "case", "defer",
			"go", "map", "struct", "chan", "else", "goto", "package", "switch",
			"const", "fallthrough", "if", "range", "type", "continue", "for",
			"import", "return", "var")
}
