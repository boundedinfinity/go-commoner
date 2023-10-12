package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/utf8"
)

func Replace[T ~string, U ~string, V ~string](s T, old U, new V) string {
	return strings.ReplaceAll(string(s), string(old), string(new))
}

func ReplaceInList[T ~string](s T, olds []string, new string) string {
	return slicer.FoldLeft(string(s), olds, func(current, old string) string {
		return Replace(current, old, new)
	})
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

func ReplaceNumbers[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf8.ToStrings(utf8.NUMBERS), replacement)
}

func ReplaceLetters[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf8.ToStrings(utf8.LETTERS), replacement)
}

func ReplaceSymbols[T ~string](s T, replacement string) string {
	return ReplaceInList(s, utf8.ToStrings(utf8.SYMBOLS), replacement)
}

func ReplaceNonNumbers[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf8.ToStrings(utf8.NUMBERS), replacement)
}

func ReplaceNonLetters[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf8.ToStrings(utf8.LETTERS), replacement)
}

func ReplaceNonWord[T ~string](s T, replacement string) string {
	return replaceNotInList(s, utf8.ToStrings(utf8.WORD_CHARACTERS), replacement)
}
