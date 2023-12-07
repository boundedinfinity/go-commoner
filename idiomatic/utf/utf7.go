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

var Utf7 = utf7{}

type utf7 struct{}

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
	return []UtfChar{
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
	}
}

func (t utf7) ControlCharacters() []UtfChar {
	return []UtfChar{
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
	}
}

func (t utf7) UpperCase() []UtfChar {
	return []UtfChar{
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
	}
}

func (t utf7) LowerCase() []UtfChar {
	return []UtfChar{
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
	}
}

func (t utf7) Letters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.UpperCase()...)
	letters = append(letters, t.LowerCase()...)
	return letters
}

func (t utf7) WordCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.Letters()...)
	letters = append(letters, t.Numbers()...)
	return letters
}

func (t utf7) VisibleCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.WordCharacters()...)
	letters = append(letters, t.Symbols()...)
	return letters
}

func (t utf7) PrintableCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.VisibleCharacters()...)
	letters = append(letters, t.Spaces()...)
	return letters
}

func (t utf7) Numbers() []UtfChar {
	return []UtfChar{
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
	}
}

func (t utf7) Symbols() []UtfChar {
	return []UtfChar{
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
	}
}

func (t utf7) Spaces() []UtfChar {
	return []UtfChar{
		BACKSPACE,
		HORIZONTAL_TAB,
		LINE_FEED,
		VERTICAL_TAB,
		FORM_FEED,
		CARRIAGE_RETURN,
		SPACE,
	}
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
