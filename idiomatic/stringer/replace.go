package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utf"
)

func Replace[T ~string, U ~string, V ~string](s T, old U, new V) string {
	return strings.ReplaceAll(string(s), string(old), string(new))
}

func ReplaceInList[T ~string](s T, olds []string, new string) string {
	return slicer.FoldLeft(string(s), func(_ int, current, old string) string {
		return Replace(current, old, new)
	}, olds...)
}

func replaceNotInList[T ~string](s T, list []string, replacement string) string {
	var n string

	for _, r := range s {
		if slicer.Contains(string(r), list...) {
			n += string(r)
		} else {
			n += replacement
		}
	}

	return n
}

func ReplaceSpace[T ~string](s T, replacement string) string {
	return ReplaceInList(s, []string{" "}, replacement)
}

func ReplaceSpaces[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf.Utf8.ToStrings(utf.Utf8.Spaces()), replacement)
}

func ReplaceNumbers[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf.Utf8.ToStrings(utf.Utf8.Numbers()), replacement)
}

func ReplaceLetters[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf.Utf8.ToStrings(utf.Utf8.Letters()), replacement)
}

func ReplaceSymbols[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf.Utf8.ToStrings(utf.Utf8.Symbols()), replacement)
}

func ReplaceNonNumbers[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf.Utf8.ToStrings(utf.Utf8.Numbers()), replacement)
}

func ReplaceNonLetters[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf.Utf8.ToStrings(utf.Utf8.Letters()), replacement)
}

func ReplaceNonWord[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf.Utf8.ToStrings(utf.Utf8.WordCharacters()), replacement)
}

func ReplaceNewlines[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf.Utf8.ToStrings(utf.Utf8.Newlines()), replacement)
}
