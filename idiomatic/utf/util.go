package utf

import "github.com/boundedinfinity/go-commoner/idiomatic/numberer"

func Range(start, end UtfChar) []UtfChar {
	return numberer.Range(start, end)
}
