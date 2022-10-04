package utf8

import (
	"github.com/boundedinfinity/go-commoner/slicer"
)

func ToStrings(cs []Utf8Char) []string {
	return slicer.Map(cs, func(c Utf8Char) string {
		return string(c)
	})
}

func IsUpperCase[T ~byte](v T) bool {
	return slicer.Contains(UPPERCASE, Utf8Char(v))
}

func IsLowerCase[T ~byte](v T) bool {
	return slicer.Contains(LOWERCASE, Utf8Char(v))
}

func IsLetter[T ~byte](v T) bool {
	return slicer.Contains(LETTERS, Utf8Char(v))
}

func IsNumber[T ~byte](v T) bool {
	return slicer.Contains(NUMBERS, Utf8Char(v))
}

func IsWhiteSpace[T ~byte](v T) bool {
	return slicer.Contains(WHITESPACE, Utf8Char(v))
}

func IsSymbols[T ~byte](v T) bool {
	return slicer.Contains(SYMBOLS, Utf8Char(v))
}
