package utfer

import (
	"strings"

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

func OneOf[T ~byte](v T, elems []UtfChar) bool {
	b := UtfChar(v)

	for _, elem := range elems {
		if elem == b {
			return true
		}
	}

	return false
}

func Replace[T ~string](s T, new UtfChar, olds ...UtfChar) string {
	var result strings.Builder

	for _, c := range []UtfChar(s) {
		r := c

		for _, old := range olds {
			if old == c {
				r = new
			}
		}

		result.WriteByte(byte(r))
	}

	return result.String()
}

func Keep[T ~string](s T, new UtfChar, keeps ...UtfChar) string {
	var result strings.Builder

	for _, c := range []UtfChar(s) {
		r := c

		for _, keep := range keeps {
			if keep != c {
				r = new
			}
		}

		result.WriteByte(byte(r))
	}

	return result.String()
}

func ReplaceSpaces[T ~string](s T, new byte) string {
	return Replace(s, UtfChar(new), Utf8.Spaces()...)
}

func ReplaceNumbers[T ~string](s T, new byte) string {
	return Replace(s, UtfChar(new), Utf8.Numbers()...)
}

func ReplaceLetters[T ~string](s T, new byte) string {
	return Replace(s, UtfChar(new), Utf8.Letters()...)
}

func ReplaceSymbols[T ~string](s T, new byte) string {
	return Replace(s, UtfChar(new), Utf8.Symbols()...)
}

func ReplaceNewlines[T ~string](s T, new byte) string {
	return Replace(s, UtfChar(new), Utf8.Newlines()...)
}

func ReplaceNonNumbers[T ~string](s T, new byte) string {
	return Keep(s, UtfChar(new), Utf8.Numbers()...)
}

func ReplaceNonLetters[T ~string](s T, new byte) string {
	return Keep(s, UtfChar(new), Utf8.Letters()...)
}

func ReplaceNonWord[T ~string](s T, new byte) string {
	return Keep(s, UtfChar(new), Utf8.WordCharacters()...)
}

func Remove[T ~string](s T, olds ...UtfChar) string {
	var result strings.Builder

	for _, c := range []UtfChar(s) {
		var found bool

		for _, old := range olds {
			if old == c {
				found = true
				break
			}
		}

		if !found {
			result.WriteByte(byte(c))
		}

	}

	return result.String()
}

func RemoveNumbers[T ~string](s T) string {
	return Remove(s, Utf8.Numbers()...)
}

func RemoveLetters[T ~string](s T) string {
	return Remove(s, Utf8.Letters()...)
}

func RemoveSymbols[T ~string](s T) string {
	return Remove(s, Utf8.Symbols()...)
}

func RemoveSpaces[T ~string](s T) string {
	return Remove(s, Utf8.Spaces()...)
}

func RemoveNewlines[T ~string](s T) string {
	return Remove(s, Utf8.Newlines()...)
}
