package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/slices"
	"github.com/boundedinfinity/go-commoner/utf8"
)

func Replace[T ~string](s T, old, new string) string {
	return strings.ReplaceAll(string(s), old, new)
}

func ReplaceInList[T ~string](s T, olds []string, new string) string {
	return slices.Fold(string(s), olds, func(current, old string) string {
		return Replace(current, old, new)
	})
}

func replaceNotInList[T ~string](s T, list []string, replacement string) string {
	var n string

	for _, r := range s {
		if slices.Contains(list, string(r)) {
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
