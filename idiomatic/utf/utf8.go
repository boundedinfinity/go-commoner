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

var Utf8 = utf8{
	utf8Toutf7:          map[UtfChar]UtfChar{},
	upperCase:           []UtfChar{},
	lowerCase:           []UtfChar{},
	extendedCharacters:  []UtfChar{},
	letters:             []UtfChar{},
	wordCharacters:      []UtfChar{},
	visibleCharacters:   []UtfChar{},
	numbers:             []UtfChar{},
	printableCharacters: []UtfChar{},
	spaces:              []UtfChar{},
	unused:              []UtfChar{},
	symbols:             []UtfChar{},
	newlines:            []UtfChar{},
	all:                 []UtfChar{},
}

func init() {
	Utf8.lowerCase = append(Utf8.lowerCase, Utf7.LowerCase()...)
	Utf8.lowerCase = append(Utf8.lowerCase,
		LOWERCASE_F_WITH_HOOK,
		LOWERCASE_LIGATURE_OE,
		LOWERCASE_Z_WITH_CARON,
		LOWERCASE_SHARP_S,
		LOWERCASE_A_WITH_GRAVE,
		LOWERCASE_A_WITH_ACUTE,
		LOWERCASE_A_WITH_CIRCUMFLEX,
		LOWERCASE_A_WITH_TILDE,
		LOWERCASE_A_WITH_DIAERESIS,
		LOWERCASE_A_WITH_RING_ABOVE,
		LOWERCASE_AE,
		LOWERCASE_C_WITH_CEDILLA,
		LOWERCASE_E_WITH_GRAVE,
		LOWERCASE_E_WITH_ACUTE,
		LOWERCASE_E_WITH_CIRCUMFLEX,
		LOWERCASE_E_WITH_DIAERESIS,
		LOWERCASE_I_WITH_GRAVE,
		LOWERCASE_I_WITH_ACUTE,
		LOWERCASE_I_WITH_CIRCUMFLEX,
		LOWERCASE_I_WITH_DIAERESIS,
		LOWERCASE_ETH,
		LOWERCASE_N_WITH_TILDE,
		LOWERCASE_O_WITH_GRAVE,
		LOWERCASE_O_WITH_ACUTE,
		LOWERCASE_O_WITH_CIRCUMFLEX,
		LOWERCASE_O_WITH_TILDE,
		LOWERCASE_O_WITH_DIAERESIS,
		LOWERCASE_O_WITH_SLASH,
		LOWERCASE_U_WITH_GRAVE,
		LOWERCASE_U_WITH_ACUTE,
		LOWERCASE_U_WITH_CIRCUMFLEX,
		LOWERCASE_U_WITH_DIAERESIS,
		LOWERCASE_Y_WITH_ACUTE,
		LOWERCASE_THORN,
		LOWERCASE_Y_WITH_DIAERESIS,
	)

	Utf8.upperCase = append(Utf8.upperCase, Utf7.UpperCase()...)
	Utf8.upperCase = append(Utf8.upperCase,
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
		UPPERCASE_E_WITH_GRAVE,
		UPPERCASE_E_WITH_ACUTE,
		UPPERCASE_E_WITH_CIRCUMFLEX,
		UPPERCASE_E_WITH_DIAERESIS,
		UPPERCASE_I_WITH_GRAVE,
		UPPERCASE_I_WITH_ACUTE,
		UPPERCASE_I_WITH_CIRCUMFLEX,
		UPPERCASE_I_WITH_DIAERESIS,
		UPPERCASE_ETH,
		UPPERCASE_N_WITH_TILDE,
		UPPERCASE_O_WITH_GRAVE,
		UPPERCASE_O_WITH_ACUTE,
		UPPERCASE_O_WITH_CIRCUMFLEX,
		UPPERCASE_O_WITH_TILDE,
		UPPERCASE_O_WITH_DIAERESIS,
		UPPERCASE_O_WITH_SLASH,
		UPPERCASE_U_WITH_GRAVE,
		UPPERCASE_U_WITH_ACUTE,
		UPPERCASE_U_WITH_CIRCUMFLEX,
		UPPERCASE_U_WITH_DIAERESIS,
		UPPERCASE_Y_WITH_ACUTE,
		UPPERCASE_THORN,
	)

	Utf8.extendedCharacters = append(Utf8.extendedCharacters,
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
		UPPERCASE_E_WITH_GRAVE,
		UPPERCASE_E_WITH_ACUTE,
		UPPERCASE_E_WITH_CIRCUMFLEX,
		UPPERCASE_E_WITH_DIAERESIS,
		UPPERCASE_I_WITH_GRAVE,
		UPPERCASE_I_WITH_ACUTE,
		UPPERCASE_I_WITH_CIRCUMFLEX,
		UPPERCASE_I_WITH_DIAERESIS,
		UPPERCASE_ETH,
		UPPERCASE_N_WITH_TILDE,
		UPPERCASE_O_WITH_GRAVE,
		UPPERCASE_O_WITH_ACUTE,
		UPPERCASE_O_WITH_CIRCUMFLEX,
		UPPERCASE_O_WITH_TILDE,
		UPPERCASE_O_WITH_DIAERESIS,
		MULTIPLICATION_SIGN,
		UPPERCASE_O_WITH_SLASH,
		UPPERCASE_U_WITH_GRAVE,
		UPPERCASE_U_WITH_ACUTE,
		UPPERCASE_U_WITH_CIRCUMFLEX,
		UPPERCASE_U_WITH_DIAERESIS,
		UPPERCASE_Y_WITH_ACUTE,
		UPPERCASE_THORN,
		LOWERCASE_SHARP_S,
		LOWERCASE_A_WITH_GRAVE,
		LOWERCASE_A_WITH_ACUTE,
		LOWERCASE_A_WITH_CIRCUMFLEX,
		LOWERCASE_A_WITH_TILDE,
		LOWERCASE_A_WITH_DIAERESIS,
		LOWERCASE_A_WITH_RING_ABOVE,
		LOWERCASE_AE,
		LOWERCASE_C_WITH_CEDILLA,
		LOWERCASE_E_WITH_GRAVE,
		LOWERCASE_E_WITH_ACUTE,
		LOWERCASE_E_WITH_CIRCUMFLEX,
		LOWERCASE_E_WITH_DIAERESIS,
		LOWERCASE_I_WITH_GRAVE,
		LOWERCASE_I_WITH_ACUTE,
		LOWERCASE_I_WITH_CIRCUMFLEX,
		LOWERCASE_I_WITH_DIAERESIS,
		LOWERCASE_ETH,
		LOWERCASE_N_WITH_TILDE,
		LOWERCASE_O_WITH_GRAVE,
		LOWERCASE_O_WITH_ACUTE,
		LOWERCASE_O_WITH_CIRCUMFLEX,
		LOWERCASE_O_WITH_TILDE,
		LOWERCASE_O_WITH_DIAERESIS,
		DIVISION_SIGN,
		LOWERCASE_O_WITH_SLASH,
		LOWERCASE_U_WITH_GRAVE,
		LOWERCASE_U_WITH_ACUTE,
		LOWERCASE_U_WITH_CIRCUMFLEX,
		LOWERCASE_U_WITH_DIAERESIS,
		LOWERCASE_Y_WITH_ACUTE,
		LOWERCASE_THORN,
		LOWERCASE_Y_WITH_DIAERESIS,
	)

	Utf8.numbers = append(Utf8.numbers, Utf7.Numbers()...)
	Utf8.numbers = append(Utf8.numbers,
		SUPERSCRIPT_1,
		SUPERSCRIPT_2_SQUARED,
		SUPERSCRIPT_3_CUBED,
		FRACTION_1_4,
		FRACTION_1_2,
		FRACTION_3_4,
	)

	Utf8.spaces = append(Utf8.spaces, Utf7.Spaces()...)
	Utf8.spaces = append(Utf8.spaces, NON_BREAKING_SPACE)

	Utf8.unused = append(Utf8.unused,
		UNUSED_1,
		UNUSED_2,
		UNUSED_3,
		UNUSED_4,
		UNUSED_5,
	)

	Utf8.symbols = append(Utf8.symbols, Utf7.Symbols()...)
	Utf8.symbols = append(Utf8.symbols,
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
	)

	Utf8.letters = append(Utf8.letters, Utf8.UpperCase()...)
	Utf8.letters = append(Utf8.letters, Utf8.LowerCase()...)

	Utf8.wordCharacters = append(Utf8.wordCharacters, Utf8.Letters()...)
	Utf8.wordCharacters = append(Utf8.wordCharacters, Utf8.Numbers()...)

	Utf8.visibleCharacters = append(Utf8.visibleCharacters, Utf8.WordCharacters()...)
	Utf8.visibleCharacters = append(Utf8.visibleCharacters, Utf8.Symbols()...)

	Utf8.printableCharacters = append(Utf8.printableCharacters, Utf8.VisibleCharacters()...)
	Utf8.printableCharacters = append(Utf8.printableCharacters, Utf8.Spaces()...)

	Utf8.all = append(Utf8.all, Utf7.All()...)
	Utf8.all = append(Utf8.all, Utf8.ExtendedCharacters()...)

	Utf8.newlines = append(Utf8.newlines, Utf7.Newlines()...)

	Utf8.utf8Toutf7 = map[UtfChar]UtfChar{
		UPPERCASE_A_WITH_GRAVE:      UPPERCASE_A,
		UPPERCASE_A_WITH_ACUTE:      UPPERCASE_A,
		UPPERCASE_A_WITH_CIRCUMFLEX: UPPERCASE_A,
		UPPERCASE_A_WITH_TILDE:      UPPERCASE_A,
		UPPERCASE_A_WITH_DIARESIS:   UPPERCASE_A,
		UPPERCASE_A_WITH_RING_ABOVE: UPPERCASE_A,
		UPPERCASE_C_WITH_CEDILLA:    UPPERCASE_E,
		UPPERCASE_E_WITH_GRAVE:      UPPERCASE_E,
		UPPERCASE_E_WITH_ACUTE:      UPPERCASE_E,
		UPPERCASE_E_WITH_CIRCUMFLEX: UPPERCASE_E,
		UPPERCASE_E_WITH_DIAERESIS:  UPPERCASE_E,
		UPPERCASE_I_WITH_GRAVE:      UPPERCASE_I,
		UPPERCASE_I_WITH_ACUTE:      UPPERCASE_I,
		UPPERCASE_I_WITH_CIRCUMFLEX: UPPERCASE_I,
		UPPERCASE_I_WITH_DIAERESIS:  UPPERCASE_I,
		UPPERCASE_N_WITH_TILDE:      UPPERCASE_N,
		UPPERCASE_O_WITH_GRAVE:      UPPERCASE_O,
		UPPERCASE_O_WITH_ACUTE:      UPPERCASE_O,
		UPPERCASE_O_WITH_CIRCUMFLEX: UPPERCASE_O,
		UPPERCASE_O_WITH_TILDE:      UPPERCASE_O,
		UPPERCASE_O_WITH_DIAERESIS:  UPPERCASE_O,
		UPPERCASE_O_WITH_SLASH:      UPPERCASE_O,
		UPPERCASE_U_WITH_GRAVE:      UPPERCASE_U,
		UPPERCASE_U_WITH_ACUTE:      UPPERCASE_U,
		UPPERCASE_U_WITH_CIRCUMFLEX: UPPERCASE_U,
		UPPERCASE_U_WITH_DIAERESIS:  UPPERCASE_U,
		UPPERCASE_S_WITH_CARON:      UPPERCASE_S,
		UPPERCASE_Y_WITH_DIAERESIS:  UPPERCASE_Y,
		UPPERCASE_Y_WITH_ACUTE:      UPPERCASE_Y,
		UPPERCASE_Z_WITH_CARON:      UPPERCASE_Z,
		LOWERCASE_A_WITH_GRAVE:      LOWERCASE_A,
		LOWERCASE_A_WITH_ACUTE:      LOWERCASE_A,
		LOWERCASE_A_WITH_CIRCUMFLEX: LOWERCASE_A,
		LOWERCASE_A_WITH_TILDE:      LOWERCASE_A,
		LOWERCASE_A_WITH_DIAERESIS:  LOWERCASE_A,
		LOWERCASE_A_WITH_RING_ABOVE: LOWERCASE_A,
		LOWERCASE_C_WITH_CEDILLA:    LOWERCASE_C,
		LOWERCASE_E_WITH_GRAVE:      LOWERCASE_E,
		LOWERCASE_E_WITH_ACUTE:      LOWERCASE_E,
		LOWERCASE_E_WITH_CIRCUMFLEX: LOWERCASE_E,
		LOWERCASE_E_WITH_DIAERESIS:  LOWERCASE_E,
		LOWERCASE_F_WITH_HOOK:       LOWERCASE_F,
		LOWERCASE_I_WITH_GRAVE:      LOWERCASE_I,
		LOWERCASE_I_WITH_ACUTE:      LOWERCASE_I,
		LOWERCASE_I_WITH_CIRCUMFLEX: LOWERCASE_I,
		LOWERCASE_I_WITH_DIAERESIS:  LOWERCASE_I,
		LOWERCASE_N_WITH_TILDE:      LOWERCASE_N,
		LOWERCASE_O_WITH_GRAVE:      LOWERCASE_O,
		LOWERCASE_O_WITH_ACUTE:      LOWERCASE_O,
		LOWERCASE_O_WITH_CIRCUMFLEX: LOWERCASE_O,
		LOWERCASE_O_WITH_TILDE:      LOWERCASE_O,
		LOWERCASE_O_WITH_DIAERESIS:  LOWERCASE_O,
		LOWERCASE_O_WITH_SLASH:      LOWERCASE_O,
		LOWERCASE_SHARP_S:           LOWERCASE_S,
		LOWERCASE_U_WITH_GRAVE:      LOWERCASE_U,
		LOWERCASE_U_WITH_ACUTE:      LOWERCASE_U,
		LOWERCASE_U_WITH_CIRCUMFLEX: LOWERCASE_U,
		LOWERCASE_U_WITH_DIAERESIS:  LOWERCASE_U,
		LOWERCASE_Y_WITH_ACUTE:      LOWERCASE_Y,
		LOWERCASE_Y_WITH_DIAERESIS:  LOWERCASE_Y,
		LOWERCASE_Z_WITH_CARON:      LOWERCASE_Z,
	}
}

type utf8 struct {
	utf8Toutf7          map[UtfChar]UtfChar
	extendedCharacters  []UtfChar
	upperCase           []UtfChar
	lowerCase           []UtfChar
	letters             []UtfChar
	wordCharacters      []UtfChar
	visibleCharacters   []UtfChar
	numbers             []UtfChar
	printableCharacters []UtfChar
	spaces              []UtfChar
	unused              []UtfChar
	symbols             []UtfChar
	newlines            []UtfChar
	all                 []UtfChar
}

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

func (t utf8) IsExtend(s byte) bool {
	return slicer.ContainsFn(func(v UtfChar) bool {
		return byte(v) == s
	}, t.ExtendedCharacters()...)
}

func (t utf8) ExtendedCharacters() []UtfChar {
	return t.extendedCharacters
}

func (t utf8) ToUtf7(fillChar UtfChar, chars ...UtfChar) []UtfChar {
	var utf7 []UtfChar

	for _, char := range chars {
		if found, ok := t.utf8Toutf7[char]; ok {
			utf7 = append(utf7, found)
		} else {
			utf7 = append(utf7, fillChar)
		}
	}

	return utf7
}

func (t utf8) All() []UtfChar {
	return t.all
}

func (t utf8) UpperCase() []UtfChar {
	return t.upperCase
}

func (t utf8) LowerCase() []UtfChar {
	return t.lowerCase
}

func (t utf8) Letters() []UtfChar {
	return t.letters
}

func (t utf8) WordCharacters() []UtfChar {
	return t.wordCharacters
}

func (t utf8) VisibleCharacters() []UtfChar {
	return t.visibleCharacters
}

func (t utf8) PrintableCharacters() []UtfChar {
	return t.printableCharacters
}

func (t utf8) Numbers() []UtfChar {
	return t.numbers
}

func (t utf8) Unused() []UtfChar {
	return t.unused
}

func (t utf8) Symbols() []UtfChar {
	return t.symbols
}

func (t utf8) Spaces() []UtfChar {
	return t.spaces
}

func (t utf8) Newlines() []UtfChar {
	return t.newlines
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
