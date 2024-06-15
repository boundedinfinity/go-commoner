package utfer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/numberer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func InRange(c, start, end UtfChar) bool {
	return c > start && c < end
}

func Range(start, end UtfChar) []UtfChar {
	return numberer.Range(start, end)
}

func ToStrings(cs []UtfChar) []string {
	return slicer.Map(func(_ int, c UtfChar) string {
		return string(c)
	}, cs...)
}
