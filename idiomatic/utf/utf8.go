package utf

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

var (
	ErrUtf8Invalid  = errors.New("invalid utf-8 character")
	ErrUtf8Invalidv = errorer.Wrapv(ErrUtf8Invalid)
)

var Utf8 = utf8{}

type utf8 struct{}

func (t utf8) Range(start, end UtfChar) []UtfChar {
	return Range(start, end)
}

func (t utf8) Parse(v byte) (UtfChar, error) {
	f, ok := slicer.FindFn(func(x UtfChar) bool {
		return UtfChar(v) == x
	}, t.All()...)

	if !ok {
		return f, ErrUtf8Invalidv(v)
	}

	return f, nil
}

func (t utf8) Is(s byte) bool {
	return slicer.ContainsFn(func(v UtfChar) bool {
		return byte(v) == s
	}, t.All()...)
}

func (t utf8) ExtendedCharacters() []UtfChar {
	return []UtfChar{
		EURO_SIGN,
		UNUSED_1,
		SINGLE_LOW_9_QUOTATION_MARK,
		LOWERCASE_F_WITH_HOOK,
		DOUBLE_LOW_9_QUOTATION_MARK,
		HORIZONTAL_ELLIPSIS,
		DAGGER,
		DOUBLE_DAGGER,
		MODIFIER_LETTER_CIRCUMFLEX_ACCENT,
		PER_MILE_SIGN,
		UPPERCASE_S_WITH_CARON,
		SINGLE_LEFT_POINTING_ANGLE_QUOTATION,
		LATIN_CAPTIAL_LIGATURE_OE,
		UNUSED_2,
		UPPERCASE_Z_WITH_CARON,
		UNUSED_3,
		UNUSED_4,
		LEFT_SINGLE_QUOTATION_MARK,
		RIGHT_SINGLE_QUOTATION_MARK,
		LEFT_DOUBLE_QUOTATION_MARK,
		RIGHT_DOUBLE_QUOTATION_MARK,
		BULLET,
		EN_DASH,
		EM_DASH,
		SMALL_TILDE,
		TRADE_MARK_SIGN,
		LATIN_SMALL_S_WITH_CARON,
		SINGLE_RIGHT_POINTING_ANBLE_QUOTATION_MARK,
		LOWERCASE_LIGATURE_OE,
		UNUSED_5,
		LOWERCASE_Z_WITH_CARON,
		UPPERCASE_Y_WITH_DIAERESIS,
		NON_BREAKING_SPACE,
		INVERTED_EXCLAMATION_MARK,
		CENT_SIGN,
		POUND_SIGN,
		CURRENCY_SIGN,
		YEN_SIGN,
		BROKEN_VERTICAL_BAR,
		SECTION_SIGN,
		SPACING_DIARESIS_UMLAUT,
		COPYRIGHT_SIGN,
		FEMININE_ORDINAL_INDICATOR,
		LEFT_DOUBLE_ANGLE_QUOTES,
		NEGATION,
		SOFT_HYPHEN,
		REGISTERED_TRADE_MARK_SIGN,
		SPACING_MACRON_OVERLINE,
		DEGREE_SIGN,
		PLUS_OR_MINUS_SIGN,
		SUPERSCRIPT_2_SQUARED,
		SUPERSCRIPT_3_CUBED,
		ACUTE_ACCENT_SPACING_ACCUTE,
		MICRO_SIGN,
		PILCROW_SIGN_PARAGRAPH_SIGN,
		MIDDLE_DOT_GEORGIAN_COMMA,
		SPACING_CEDILLA,
		SUPERSCRIPT_1,
		MASCULINE_ORDINAL_INDICATOR,
		RIGHT_DOUBLE_ANGLE_QUOTES,
		FRACTION_1_4,
		FRACTION_1_2,
		FRACTION_3_4,
		INVERTED_QUESTION_MARK,
		UPPERCASE_A_WITH_GRAVE,
		UPPERCASE_A_WITH_ACUTE,
		UPPERCASE_A_WITH_CIRCUMFLEX,
		UPPERCASE_A_WITH_TILDE,
		UPPERCASE_A_WITH_DIARESIS,
		UPPERCASE_A_WITH_RING_ABOVE,
		UPPERCASE_AE,
		UPPERCASE_C_WITH_CEDILLA,
		UPPERCASE_e_WITH_GRAVE,
		UPPERCASE_E_WITH_ACUTE,
		UPPERCASE_E_WITH_circumflex,
		UPPERCASE_E_WITH_DIAERESIS,
		UPPERCASE_I_WITH_GRAVE,
		UPPERCASE_I_WITH_ACUTE,
		UPPERCASE_I_WITH_circumflex,
		UPPERCASE_I_WITH_DIAERESIS,
		UPPERCASE_ETH,
		UPPERCASE_N_WITH_TILDE,
		UPPERCASE_O_WITH_GRAVE,
		UPPERCASE_O_WITH_ACUTE,
		UPPERCASE_O_WITH_circumflex,
		UPPERCASE_O_WITH_TILDE,
		UPPERCASE_O_WITH_DIAERESIS,
		MULTIPLICATION_SIGN,
		UPPERCASE_O_WITH_SLASH,
		UPPERCASE_U_WITH_GRAVE,
		UPPERCASE_U_WITH_ACUTE,
		UPPERCASE_U_WITH_circumflex,
		UPPERCASE_U_WITH_DIAERESIS,
		UPPERCASE_Y_WITH_ACUTE,
		UPPERCASE_THORN,
		LOWERCASE_SHARP_S,
		LOWERCASE_A_WITH_GRAVE,
		LOWERCASE_A_WITH_ACUTE,
		LOWERCASE_A_WITH_circumflex,
		LOWERCASE_A_WITH_TILDE,
		LOWERCASE_A_WITH_DIAERESIS,
		LOWERCASE_A_WITH_RING_ABOVE,
		LOWERCASE_AE,
		LOWERCASE_C_WITH_CEDILLA,
		LOWERCASE_E_WITH_GRAVE,
		LOWERCASE_E_WITH_ACUTE,
		LOWERCASE_E_WITH_circumflex,
		LOWERCASE_E_WITH_DIAERESIS,
		LOWERCASE_I_WITH_GRAVE,
		LOWERCASE_I_WITH_ACUTE,
		LOWERCASE_I_WITH_circumflex,
		LOWERCASE_I_WITH_DIAERESIS,
		LOWERCASE_ETH,
		LOWERCASE_N_WITH_TILDE,
		LOWERCASE_O_WITH_GRAVE,
		LOWERCASE_O_WITH_ACUTE,
		LOWERCASE_O_WITH_circumflex,
		LOWERCASE_O_WITH_TILDE,
		LOWERCASE_O_WITH_DIAERESIS,
		DIVISION_SIGN,
		LOWERCASE_O_WITH_SLASH,
		LOWERCASE_U_WITH_GRAVE,
		LOWERCASE_U_WITH_ACUTE,
		LOWERCASE_U_WITH_circumflex,
		LOWERCASE_U_WITH_DIAERESIS,
		LOWERCASE_Y_WITH_ACUTE,
		LOWERCASE_THORN,
		LOWERCASE_Y_WITH_DIAERESIS,
	}
}

func (t utf8) All() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.All()...)
	chars = append(chars, t.ExtendedCharacters()...)

	return chars
}

func (t utf8) UpperCase() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.UpperCase()...)
	chars = append(chars, []UtfChar{
		UPPERCASE_S_WITH_CARON,
		UPPERCASE_Z_WITH_CARON,
		UPPERCASE_Y_WITH_DIAERESIS,
		UPPERCASE_A_WITH_GRAVE,
		UPPERCASE_A_WITH_ACUTE,
		UPPERCASE_A_WITH_CIRCUMFLEX,
		UPPERCASE_A_WITH_TILDE,
		UPPERCASE_A_WITH_DIARESIS,
		UPPERCASE_A_WITH_RING_ABOVE,
		UPPERCASE_AE,
		UPPERCASE_C_WITH_CEDILLA,
		UPPERCASE_e_WITH_GRAVE,
		UPPERCASE_E_WITH_ACUTE,
		UPPERCASE_E_WITH_circumflex,
		UPPERCASE_E_WITH_DIAERESIS,
		UPPERCASE_I_WITH_GRAVE,
		UPPERCASE_I_WITH_ACUTE,
		UPPERCASE_I_WITH_circumflex,
		UPPERCASE_I_WITH_DIAERESIS,
		UPPERCASE_ETH,
		UPPERCASE_N_WITH_TILDE,
		UPPERCASE_O_WITH_GRAVE,
		UPPERCASE_O_WITH_ACUTE,
		UPPERCASE_O_WITH_circumflex,
		UPPERCASE_O_WITH_TILDE,
		UPPERCASE_O_WITH_DIAERESIS,
		UPPERCASE_O_WITH_SLASH,
		UPPERCASE_U_WITH_GRAVE,
		UPPERCASE_U_WITH_ACUTE,
		UPPERCASE_U_WITH_circumflex,
		UPPERCASE_U_WITH_DIAERESIS,
		UPPERCASE_Y_WITH_ACUTE,
		UPPERCASE_THORN,
	}...)

	return chars
}

func (t utf8) LowerCase() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.LowerCase()...)
	chars = append(chars, []UtfChar{
		LOWERCASE_F_WITH_HOOK,
		LOWERCASE_LIGATURE_OE,
		LOWERCASE_Z_WITH_CARON,
		LOWERCASE_SHARP_S,
		LOWERCASE_A_WITH_GRAVE,
		LOWERCASE_A_WITH_ACUTE,
		LOWERCASE_A_WITH_circumflex,
		LOWERCASE_A_WITH_TILDE,
		LOWERCASE_A_WITH_DIAERESIS,
		LOWERCASE_A_WITH_RING_ABOVE,
		LOWERCASE_AE,
		LOWERCASE_C_WITH_CEDILLA,
		LOWERCASE_E_WITH_GRAVE,
		LOWERCASE_E_WITH_ACUTE,
		LOWERCASE_E_WITH_circumflex,
		LOWERCASE_E_WITH_DIAERESIS,
		LOWERCASE_I_WITH_GRAVE,
		LOWERCASE_I_WITH_ACUTE,
		LOWERCASE_I_WITH_circumflex,
		LOWERCASE_I_WITH_DIAERESIS,
		LOWERCASE_ETH,
		LOWERCASE_N_WITH_TILDE,
		LOWERCASE_O_WITH_GRAVE,
		LOWERCASE_O_WITH_ACUTE,
		LOWERCASE_O_WITH_circumflex,
		LOWERCASE_O_WITH_TILDE,
		LOWERCASE_O_WITH_DIAERESIS,
		LOWERCASE_O_WITH_SLASH,
		LOWERCASE_U_WITH_GRAVE,
		LOWERCASE_U_WITH_ACUTE,
		LOWERCASE_U_WITH_circumflex,
		LOWERCASE_U_WITH_DIAERESIS,
		LOWERCASE_Y_WITH_ACUTE,
		LOWERCASE_THORN,
		LOWERCASE_Y_WITH_DIAERESIS,
	}...)

	return chars
}

func (t utf8) Letters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.UpperCase()...)
	letters = append(letters, t.LowerCase()...)
	return letters
}

func (t utf8) WordCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.Letters()...)
	letters = append(letters, t.Numbers()...)
	return letters
}

func (t utf8) VisibleCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.WordCharacters()...)
	letters = append(letters, t.Symbols()...)
	return letters
}

func (t utf8) PrintableCharacters() []UtfChar {
	var letters []UtfChar
	letters = append(letters, t.VisibleCharacters()...)
	letters = append(letters, t.Spaces()...)
	return letters
}

func (t utf8) Numbers() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.Numbers()...)
	chars = append(chars, []UtfChar{
		SUPERSCRIPT_1,
		SUPERSCRIPT_2_SQUARED,
		SUPERSCRIPT_3_CUBED,
		FRACTION_1_4,
		FRACTION_1_2,
		FRACTION_3_4,
	}...)

	return chars
}

func (t utf8) Unused() []UtfChar {
	return []UtfChar{
		UNUSED_1,
		UNUSED_2,
		UNUSED_3,
		UNUSED_4,
		UNUSED_5,
	}
}

func (t utf8) Symbols() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.Symbols()...)
	chars = append(chars, []UtfChar{
		EURO_SIGN,
		SINGLE_LOW_9_QUOTATION_MARK,
		DOUBLE_LOW_9_QUOTATION_MARK,
		HORIZONTAL_ELLIPSIS,
		DAGGER,
		DOUBLE_DAGGER,
		MODIFIER_LETTER_CIRCUMFLEX_ACCENT,
		PER_MILE_SIGN,
		SINGLE_LEFT_POINTING_ANGLE_QUOTATION,
		LEFT_SINGLE_QUOTATION_MARK,
		RIGHT_SINGLE_QUOTATION_MARK,
		LEFT_DOUBLE_QUOTATION_MARK,
		RIGHT_DOUBLE_QUOTATION_MARK,
		BULLET,
		EN_DASH,
		EM_DASH,
		SMALL_TILDE,
		TRADE_MARK_SIGN,
		LATIN_SMALL_S_WITH_CARON,
		SINGLE_RIGHT_POINTING_ANBLE_QUOTATION_MARK,
		INVERTED_EXCLAMATION_MARK,
		CENT_SIGN,
		POUND_SIGN,
		CURRENCY_SIGN,
		YEN_SIGN,
		BROKEN_VERTICAL_BAR,
		SECTION_SIGN,
		SPACING_DIARESIS_UMLAUT,
		COPYRIGHT_SIGN,
		FEMININE_ORDINAL_INDICATOR,
		LEFT_DOUBLE_ANGLE_QUOTES,
		NEGATION,
		SOFT_HYPHEN,
		REGISTERED_TRADE_MARK_SIGN,
		SPACING_MACRON_OVERLINE,
		DEGREE_SIGN,
		PLUS_OR_MINUS_SIGN,
		ACUTE_ACCENT_SPACING_ACCUTE,
		MICRO_SIGN,
		PILCROW_SIGN_PARAGRAPH_SIGN,
		MIDDLE_DOT_GEORGIAN_COMMA,
		SPACING_CEDILLA,
		MASCULINE_ORDINAL_INDICATOR,
		RIGHT_DOUBLE_ANGLE_QUOTES,
		INVERTED_QUESTION_MARK,
		MULTIPLICATION_SIGN,
		DIVISION_SIGN,
	}...)

	return chars
}

func (t utf8) Spaces() []UtfChar {
	var chars []UtfChar
	chars = append(chars, Utf7.Spaces()...)
	chars = append(chars, NON_BREAKING_SPACE)
	return chars

}

func (t utf8) ToStrings(cs []UtfChar) []string {
	return slicer.Map(func(c UtfChar) string {
		return string(c)
	}, cs...)
}

func (t utf8) IsUpperCase(v byte) bool {
	return slicer.Contains(UtfChar(v), t.UpperCase()...)
}

func (t utf8) IsLowerCase(v byte) bool {
	return slicer.Contains(UtfChar(v), t.LowerCase()...)
}

func (t utf8) IsLetter(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Letters()...)
}

func (t utf8) IsNumber(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Numbers()...)
}

func (t utf8) IsSpace(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Spaces()...)
}

func (t utf8) IsSymbol(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Symbols()...)
}

func (t utf8) IsUnused(v byte) bool {
	return slicer.Contains(UtfChar(v), t.Unused()...)
}
