package stringer

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/runer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
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
	UppercaseLetters = append(UppercaseLetters, runer.Range(runer.UPPERCASE_A, runer.UPPERCASE_Z)...)
	LowercaseLetters = append(LowercaseLetters, runer.Range(runer.LOWERCASE_A, runer.LOWERCASE_Z)...)

	Letters = append(Letters, LowercaseLetters...)
	Letters = append(Letters, UppercaseLetters...)

	Numbers = append(Numbers, runer.Range(runer.ZERO, runer.NINE)...)

	LettersAndNumbers = append(LettersAndNumbers, Letters...)
	LettersAndNumbers = append(LettersAndNumbers, Numbers...)

	Symbols = append(Symbols, runer.Range(runer.EXCLAMATION_MARK, runer.SLASH)...)
	Symbols = append(Symbols, runer.Range(runer.COLON, runer.AT_SIGN)...)
	Symbols = append(Symbols, runer.Range(runer.LEFT_SQUARE_BRACKET, runer.GRAVE)...)
	Symbols = append(Symbols, runer.Range(runer.LEFT_CURLY_BRACKET, runer.TILDE)...)

	VisibleCharacters = append(VisibleCharacters, LettersAndNumbers...)
	VisibleCharacters = append(VisibleCharacters, Symbols...)
}

func isFn(s string, rs []rune) bool {
	for _, c := range s {
		if !slicer.Contains(c, rs...) {
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
