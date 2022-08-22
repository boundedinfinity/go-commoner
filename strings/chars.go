package strings

import (
	"github.com/boundedinfinity/go-commoner/runes"
	"github.com/boundedinfinity/go-commoner/slices"
)

var (
	UppercaseLetters  = []rune{}
	LowercaseLetters  = []rune{}
	Letters           = []rune{}
	Numbers           = []rune{}
	LettersAndNumbers = []rune{}
	Symbols           = []rune{}
	VisibleCharacters = []rune{}
)

func init() {
	UppercaseLetters = append(UppercaseLetters, runes.Range(runes.UPPERCASE_A, runes.UPPERCASE_Z)...)
	LowercaseLetters = append(LowercaseLetters, runes.Range(runes.LOWERCASE_A, runes.LOWERCASE_Z)...)

	Letters = append(Letters, LowercaseLetters...)
	Letters = append(Letters, UppercaseLetters...)

	Numbers = append(Numbers, runes.Range(runes.ZERO, runes.NINE)...)

	LettersAndNumbers = append(LettersAndNumbers, Letters...)
	LettersAndNumbers = append(LettersAndNumbers, Numbers...)

	Symbols = append(Symbols, runes.Range(runes.EXCLAMATION_MARK, runes.SLASH)...)
	Symbols = append(Symbols, runes.Range(runes.COLON, runes.AT_SIGN)...)
	Symbols = append(Symbols, runes.Range(runes.LEFT_SQUARE_BRACKET, runes.GRAVE)...)
	Symbols = append(Symbols, runes.Range(runes.LEFT_CURLY_BRACKET, runes.TILDE)...)

	VisibleCharacters = append(VisibleCharacters, LettersAndNumbers...)
	VisibleCharacters = append(VisibleCharacters, Symbols...)
}

func isFn(s string, rs []rune) bool {
	for _, c := range s {
		if !slices.Contains(rs, c) {
			return false
		}
	}

	return true
}

func IsUppercaseLetters(s string) bool {
	return isFn(s, UppercaseLetters)
}

func IsLowercaseLetters(s string) bool {
	return isFn(s, LowercaseLetters)
}

func IsLetters(s string) bool {
	return isFn(s, Letters)
}

func IsNumbers(s string) bool {
	return isFn(s, Numbers)
}

func IsLettersAndNumbers(s string) bool {
	return isFn(s, LettersAndNumbers)
}

func IsVisibleCharacters(s string) bool {
	return isFn(s, VisibleCharacters)
}
