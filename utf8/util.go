package utf8

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func ToStrings(cs []Utf8Char) []string {
	return slicer.Map(cs, func(c Utf8Char) string {
		return string(c)
	})
}

func IsUpperCase[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), UPPERCASE...)
}

func IsLowerCase[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), LOWERCASE...)
}

func IsLetter[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), LETTERS...)
}

func IsNumber[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), NUMBERS...)
}

func IsWhiteSpace[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), WHITESPACE...)
}

func IsSymbols[T ~byte](v T) bool {
	return slicer.Contains(Utf8Char(v), SYMBOLS...)
}
