package langer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
)

var (
	languageCharacters = []utfer.UtfChar{}
)

func init() {
	languageCharacters = append(languageCharacters, utfer.Utf7.WordCharacters()...)
	languageCharacters = append(languageCharacters, utfer.UNDERSCORE)
}

func ReplaceNonLanguageCharacters[T ~string](s T, replacement T, startsWithNumber T) string {
	o := replaceNotInList(s, utfer.ToStrings(languageCharacters), string(replacement))

	if len(o) > 0 && utfer.Utf7.IsNumber(o[0]) {
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
