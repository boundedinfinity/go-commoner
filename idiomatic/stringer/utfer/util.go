package utfer

import (
	"strings"
)

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
