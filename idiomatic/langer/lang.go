package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utf"
)

var (
	languageCharacters = []utf.UtfChar{}
)

func init() {
	languageCharacters = append(languageCharacters, utf.Utf7.WordCharacters()...)
	languageCharacters = append(languageCharacters, utf.UNDERSCORE)
}

func ReplaceNonLanguageCharacters[T ~string](s T, replacement T, startsWithNumber T) string {
	o := replaceNotInList(s, utf.ToStrings(languageCharacters), string(replacement))

	if len(o) > 0 && utf.Utf7.IsNumber(o[0]) {
		o = string(startsWithNumber) + o
	}

	return o
}

func RemoveNonLanguageCharacters[T ~string](s T, startsWithNumber T) string {
	return ReplaceNonLanguageCharacters(s, "", startsWithNumber)
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
