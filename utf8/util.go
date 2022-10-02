package utf8

import "github.com/boundedinfinity/go-commoner/slices"

func ToStrings(cs []Utf8Char) []string {
	return slices.Map(cs, func(c Utf8Char) string {
		return string(c)
	})
}

func IsUpperCase[T ~byte](v T) bool {
	return slices.Contains(UPPERCASE, Utf8Char(v))
}

func IsLowerCase[T ~byte](v T) bool {
	return slices.Contains(LOWERCASE, Utf8Char(v))
}

func IsLetter[T ~byte](v T) bool {
	return slices.Contains(LETTERS, Utf8Char(v))
}

func IsNumber[T ~byte](v T) bool {
	return slices.Contains(NUMBERS, Utf8Char(v))
}

func IsWhiteSpace[T ~byte](v T) bool {
	return slices.Contains(WHITESPACE, Utf8Char(v))
}

func IsSymbols[T ~byte](v T) bool {
	return slices.Contains(SYMBOLS, Utf8Char(v))
}
