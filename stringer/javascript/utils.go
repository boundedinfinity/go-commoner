package javascript

import "github.com/boundedinfinity/go-commoner/slicer"

func IsKeyword(s string) bool {
	return slicer.Contains(Keywords, s)
}
