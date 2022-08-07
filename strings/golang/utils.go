package golang

import "github.com/boundedinfinity/commons/slices"

func IsKeyword(s string) bool {
	return slices.Contains(Keywords, s)
}
