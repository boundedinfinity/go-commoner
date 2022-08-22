package javascript

import "github.com/boundedinfinity/go-commoner/slices"

func IsKeyword(s string) bool {
	return slices.Contains(Keywords, s)
}
