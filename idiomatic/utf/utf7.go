package utf

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

var (
	ErrUtf7Invalid  = errors.New("invalid utf-7 character")
	ErrUtf7Invalidv = errorer.Wrapv(ErrUtf7Invalid)
)

var Utf7 = utf7{
	all:                 []UtfChar{},
	controlCharacters:   []UtfChar{},
	upperCase:           []UtfChar{},
	lowerCase:           []UtfChar{},
	numbers:             []UtfChar{},
	letters:             []UtfChar{},
	wordCharacters:      []UtfChar{},
	visibleCharacters:   []UtfChar{},
	printableCharacters: []UtfChar{},
	symbols:             []UtfChar{},
	spaces:              []UtfChar{},
	newlines:            []UtfChar{},
}

type utf7 struct {
	all                 []UtfChar
	controlCharacters   []UtfChar
	upperCase           []UtfChar
	lowerCase           []UtfChar
	numbers             []UtfChar
	letters             []UtfChar
	wordCharacters      []UtfChar
	visibleCharacters   []UtfChar
	printableCharacters []UtfChar
	symbols             []UtfChar
	spaces              []UtfChar
	newlines            []UtfChar
}

func init() {
	Utf7.numbers = append(Utf7.numbers,
		NUMBER_ZERO,
		NUMBER_ONE,
		NUMBER_TWO,
		NUMBER_THREE,
		NUMBER_FOUR,
		NUMBER_FIVE,
		NUMBER_SIX,
		NUMBER_SEVEN,
		NUMBER_EIGHT,
		NUMBER_NINE,
	)

	Utf7.lowerCase = append(Utf7.lowerCase,
		LOWERCASE_A,
		LOWERCASE_B,
		LOWERCASE_C,
		LOWERCASE_D,
		LOWERCASE_E,
		LOWERCASE_F,
		LOWERCASE_G,
		LOWERCASE_H,
		LOWERCASE_I,
		LOWERCASE_J,
		LOWERCASE_K,
		LOWERCASE_L,
		LOWERCASE_M,
		LOWERCASE_N,
		LOWERCASE_O,
		LOWERCASE_P,
		LOWERCASE_Q,
		LOWERCASE_R,
		LOWERCASE_S,
		LOWERCASE_T,
		LOWERCASE_U,
		LOWERCASE_V,
		LOWERCASE_W,
		LOWERCASE_X,
		LOWERCASE_Y,
		LOWERCASE_Z,
	)

	Utf7.upperCase = append(Utf7.upperCase,
		UPPERCASE_A,
		UPPERCASE_B,
		UPPERCASE_C,
		UPPERCASE_D,
		UPPERCASE_E,
		UPPERCASE_F,
		UPPERCASE_G,
		UPPERCASE_H,
		UPPERCASE_I,
		UPPERCASE_J,
		UPPERCASE_K,
		UPPERCASE_L,
		UPPERCASE_M,
		UPPERCASE_N,
		UPPERCASE_O,
		UPPERCASE_P,
		UPPERCASE_Q,
		UPPERCASE_R,
		UPPERCASE_S,
		UPPERCASE_T,
		UPPERCASE_U,
		UPPERCASE_V,
		UPPERCASE_W,
		UPPERCASE_X,
		UPPERCASE_Y,
		UPPERCASE_Z,
	)

	Utf7.all = append(Utf7.all,
		NULL,
		START_OF_HEADER,
		START_OF_TEXT,
		END_OF_TEXT,
		END_OF_TRANSMISSION,
		ENQUIRY,
		ACKNOWLEDGE,
		BELL,
		BACKSPACE,
		HORIZONTAL_TAB,
		LINE_FEED,
		VERTICAL_TAB,
		FORM_FEED,
		CARRIAGE_RETURN,
		SHIFT_OUT,
		SHIFT_IN,
		DATA_LINK_ESCAPE,
		DEVICE_CONTROL_1,
		DEVICE_CONTROL_2,
		DEVICE_CONTROL_3,
		DEVICE_CONTROL_4,
		NEGATIVE_ACKNOLEDGE,
		SYNCHRONIZE,
		END_OF_TRANSMISSION_BLOCK,
		CANCEL,
		END_OF_MEDIUM,
		SUBSTITUE,
		ESCAPE,
		FILE_SEPARATOR,
		GROUP_SEPARATOR,
		RECORD_SEPARATOR,
		UNIT_SEPARATOR,
		SPACE,
		EXCLAMATION_MARK,
		DOUBLE_QUOTE,
		NUMBER_SIGN,
		DOLLAR,
		PERCENT,
		AMPERSAND,
		SINGLE_QUOTE,
		OPENING_PARENTHESIS,
		CLOSING_PARENTHESIS,
		ASTERISK,
		PLUS,
		COMMA,
		MINUS,
		PERIOD,
		SLASH,
		NUMBER_ZERO,
		NUMBER_ONE,
		NUMBER_TWO,
		NUMBER_THREE,
		NUMBER_FOUR,
		NUMBER_FIVE,
		NUMBER_SIX,
		NUMBER_SEVEN,
		NUMBER_EIGHT,
		NUMBER_NINE,
		COLON,
		SEMICOLON,
		LESS_THAN,
		EQUAL_SIGN,
		GREATER_THAN,
		QUESTION_MARK,
		AT_SIGN,
		UPPERCASE_A,
		UPPERCASE_B,
		UPPERCASE_C,
		UPPERCASE_D,
		UPPERCASE_E,
		UPPERCASE_F,
		UPPERCASE_G,
		UPPERCASE_H,
		UPPERCASE_I,
		UPPERCASE_J,
		UPPERCASE_K,
		UPPERCASE_L,
		UPPERCASE_M,
		UPPERCASE_N,
		UPPERCASE_O,
		UPPERCASE_P,
		UPPERCASE_Q,
		UPPERCASE_R,
		UPPERCASE_S,
		UPPERCASE_T,
		UPPERCASE_U,
		UPPERCASE_V,
		UPPERCASE_W,
		UPPERCASE_X,
		UPPERCASE_Y,
		UPPERCASE_Z,
		OPENING_SQUARE_BRACKET,
		BACKSLASH,
		CLOSING_SQUARE_BRACKET,
		CARET,
		UNDERSCORE,
		GRAVE_ACCENT,
		LOWERCASE_A,
		LOWERCASE_B,
		LOWERCASE_C,
		LOWERCASE_D,
		LOWERCASE_E,
		LOWERCASE_F,
		LOWERCASE_G,
		LOWERCASE_H,
		LOWERCASE_I,
		LOWERCASE_J,
		LOWERCASE_K,
		LOWERCASE_L,
		LOWERCASE_M,
		LOWERCASE_N,
		LOWERCASE_O,
		LOWERCASE_P,
		LOWERCASE_Q,
		LOWERCASE_R,
		LOWERCASE_S,
		LOWERCASE_T,
		LOWERCASE_U,
		LOWERCASE_V,
		LOWERCASE_W,
		LOWERCASE_X,
		LOWERCASE_Y,
		LOWERCASE_Z,
		OPENING_CURLY_BRACKET,
		VERTICAL_BAR,
		CLOSING_CURLY_BRACKET,
		TILDE,
		DELETE,
	)

	Utf7.controlCharacters = append(Utf7.controlCharacters,
		NULL,
		START_OF_HEADER,
		START_OF_TEXT,
		END_OF_TEXT,
		END_OF_TRANSMISSION,
		ENQUIRY,
		ACKNOWLEDGE,
		BELL,
		BACKSPACE,
		HORIZONTAL_TAB,
		LINE_FEED,
		VERTICAL_TAB,
		FORM_FEED,
		CARRIAGE_RETURN,
		SHIFT_OUT,
		SHIFT_IN,
		DATA_LINK_ESCAPE,
		DEVICE_CONTROL_1,
		DEVICE_CONTROL_2,
		DEVICE_CONTROL_3,
		DEVICE_CONTROL_4,
		NEGATIVE_ACKNOLEDGE,
		SYNCHRONIZE,
		END_OF_TRANSMISSION_BLOCK,
		CANCEL,
		END_OF_MEDIUM,
		SUBSTITUE,
		ESCAPE,
		FILE_SEPARATOR,
		GROUP_SEPARATOR,
		RECORD_SEPARATOR,
		UNIT_SEPARATOR,
	)

	Utf7.symbols = append(Utf7.symbols,
		EXCLAMATION_MARK,
		DOUBLE_QUOTE,
		NUMBER_SIGN,
		DOLLAR,
		PERCENT,
		AMPERSAND,
		SINGLE_QUOTE,
		OPENING_PARENTHESIS,
		CLOSING_PARENTHESIS,
		ASTERISK,
		PLUS,
		COMMA,
		MINUS,
		PERIOD,
		SLASH,
		COLON,
		SEMICOLON,
		LESS_THAN,
		EQUAL_SIGN,
		GREATER_THAN,
		QUESTION_MARK,
		AT_SIGN,
		OPENING_SQUARE_BRACKET,
		BACKSLASH,
		CLOSING_SQUARE_BRACKET,
		CARET,
		UNDERSCORE,
		GRAVE_ACCENT,
		OPENING_CURLY_BRACKET,
		VERTICAL_BAR,
		CLOSING_CURLY_BRACKET,
		TILDE,
	)

	Utf7.spaces = append(Utf7.spaces,
		BACKSPACE,
		HORIZONTAL_TAB,
		LINE_FEED,
		VERTICAL_TAB,
		FORM_FEED,
		CARRIAGE_RETURN,
		SPACE,
	)

	Utf7.newlines = append(Utf7.newlines,
		LINE_FEED,
		CARRIAGE_RETURN,
	)

	Utf7.letters = append(Utf7.letters, Utf7.UpperCase()...)
	Utf7.letters = append(Utf7.letters, Utf7.LowerCase()...)

	Utf7.wordCharacters = append(Utf7.wordCharacters, Utf7.Letters()...)
	Utf7.wordCharacters = append(Utf7.wordCharacters, Utf7.Numbers()...)

	Utf7.visibleCharacters = append(Utf7.visibleCharacters, Utf7.WordCharacters()...)
	Utf7.visibleCharacters = append(Utf7.visibleCharacters, Utf7.Symbols()...)

	Utf7.printableCharacters = append(Utf7.printableCharacters, Utf7.VisibleCharacters()...)
	Utf7.printableCharacters = append(Utf7.printableCharacters, Utf7.Spaces()...)
}

func (t utf7) Range(start, end UtfChar) []UtfChar {
	return Range(start, end)
}

func (t utf7) Parse(v byte) (UtfChar, error) {
	f, ok := slicer.FindFn(func(x UtfChar) bool {
		return UtfChar(v) == x
	}, t.All()...)

	if !ok {
		return f, ErrUtf7Invalidv(v)
	}

	return f, nil
}

func (t utf7) Is(s byte) bool {
	return slicer.ContainsFn(func(v UtfChar) bool {
		return byte(v) == s
	}, t.All()...)
}

func (t utf7) All() []UtfChar {
	return t.all
}

func (t utf7) ControlCharacters() []UtfChar {
	return t.controlCharacters
}

func (t utf7) UpperCase() []UtfChar {
	return t.upperCase
}

func (t utf7) LowerCase() []UtfChar {
	return t.lowerCase
}

func (t utf7) Letters() []UtfChar {
	return t.letters
}

func (t utf7) WordCharacters() []UtfChar {
	return t.wordCharacters
}

func (t utf7) VisibleCharacters() []UtfChar {
	return t.visibleCharacters
}

func (t utf7) PrintableCharacters() []UtfChar {
	return t.printableCharacters
}

func (t utf7) Numbers() []UtfChar {
	return t.numbers
}

func (t utf7) Symbols() []UtfChar {
	return t.symbols
}

func (t utf7) Spaces() []UtfChar {
	return t.spaces
}

func (t utf7) Newlines() []UtfChar {
	return t.newlines
}

func (t utf7) ToStrings(cs []UtfChar) []string {
	return slicer.Map(func(c UtfChar) string {
		return string(c)
	}, cs...)
}

func (t utf7) IsUpperCase(v byte) bool {
	return slicer.Contains(UtfChar(v), t.UpperCase()...)
}

func (t utf7) IsLowerCase(v byte) bool {
	return slicer.Contains(UtfChar(v), t.LowerCase()...)
}

func (t utf7) IsLetter(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Letters()...)
}

func (t utf7) IsNumber(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Numbers()...)
}

func (t utf7) IsSpace(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Spaces()...)
}

func (t utf7) IsSymbol(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Symbols()...)
}
