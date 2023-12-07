package utf

import (
	"encoding/json"
)

// https://www.ascii-code.com/

type UtfChar byte

func (t UtfChar) String() string {
	return string(t)
}

func (t UtfChar) MarshalJSON() ([]byte, error) {
	return json.Marshal(byte(t))
}

func (t *UtfChar) UnmarshalJSON(data []byte) error {
	var v byte

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Utf7.Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t UtfChar) MarshalYAML() (interface{}, error) {
	return byte(t), nil
}

func (t *UtfChar) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v byte

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Utf7.Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

const (
	NULL                                       UtfChar = 0
	START_OF_HEADER                            UtfChar = 1
	START_OF_TEXT                              UtfChar = 2
	END_OF_TEXT                                UtfChar = 3
	END_OF_TRANSMISSION                        UtfChar = 4
	ENQUIRY                                    UtfChar = 5
	ACKNOWLEDGE                                UtfChar = 6
	BELL                                       UtfChar = 7
	BACKSPACE                                  UtfChar = 8
	HORIZONTAL_TAB                             UtfChar = 9
	LINE_FEED                                  UtfChar = 10
	VERTICAL_TAB                               UtfChar = 11
	FORM_FEED                                  UtfChar = 12
	CARRIAGE_RETURN                            UtfChar = 13
	SHIFT_OUT                                  UtfChar = 14
	SHIFT_IN                                   UtfChar = 15
	DATA_LINK_ESCAPE                           UtfChar = 16
	DEVICE_CONTROL_1                           UtfChar = 17
	DEVICE_CONTROL_2                           UtfChar = 18
	DEVICE_CONTROL_3                           UtfChar = 19
	DEVICE_CONTROL_4                           UtfChar = 20
	NEGATIVE_ACKNOLEDGE                        UtfChar = 21
	SYNCHRONIZE                                UtfChar = 22
	END_OF_TRANSMISSION_BLOCK                  UtfChar = 23
	CANCEL                                     UtfChar = 24
	END_OF_MEDIUM                              UtfChar = 25
	SUBSTITUE                                  UtfChar = 26
	ESCAPE                                     UtfChar = 27
	FILE_SEPARATOR                             UtfChar = 28
	GROUP_SEPARATOR                            UtfChar = 29
	RECORD_SEPARATOR                           UtfChar = 30
	UNIT_SEPARATOR                             UtfChar = 31
	SPACE                                      UtfChar = 32
	EXCLAMATION_MARK                           UtfChar = 33
	DOUBLE_QUOTE                               UtfChar = 34
	NUMBER_SIGN                                UtfChar = 35
	DOLLAR                                     UtfChar = 36
	PERCENT                                    UtfChar = 37
	AMPERSAND                                  UtfChar = 38
	SINGLE_QUOTE                               UtfChar = 39
	OPENING_PARENTHESIS                        UtfChar = 40
	CLOSING_PARENTHESIS                        UtfChar = 41
	ASTERISK                                   UtfChar = 42
	PLUS                                       UtfChar = 43
	COMMA                                      UtfChar = 44
	MINUS                                      UtfChar = 45
	PERIOD                                     UtfChar = 46
	SLASH                                      UtfChar = 47
	NUMBER_ZERO                                UtfChar = 48
	NUMBER_ONE                                 UtfChar = 49
	NUMBER_TWO                                 UtfChar = 50
	NUMBER_THREE                               UtfChar = 51
	NUMBER_FOUR                                UtfChar = 52
	NUMBER_FIVE                                UtfChar = 53
	NUMBER_SIX                                 UtfChar = 54
	NUMBER_SEVEN                               UtfChar = 55
	NUMBER_EIGHT                               UtfChar = 56
	NUMBER_NINE                                UtfChar = 57
	COLON                                      UtfChar = 58
	SEMICOLON                                  UtfChar = 59
	LESS_THAN                                  UtfChar = 60
	EQUAL_SIGN                                 UtfChar = 61
	GREATER_THAN                               UtfChar = 62
	QUESTION_MARK                              UtfChar = 63
	AT_SIGN                                    UtfChar = 64
	UPPERCASE_A                                UtfChar = 65
	UPPERCASE_B                                UtfChar = 66
	UPPERCASE_C                                UtfChar = 67
	UPPERCASE_D                                UtfChar = 68
	UPPERCASE_E                                UtfChar = 69
	UPPERCASE_F                                UtfChar = 70
	UPPERCASE_G                                UtfChar = 71
	UPPERCASE_H                                UtfChar = 72
	UPPERCASE_I                                UtfChar = 73
	UPPERCASE_J                                UtfChar = 74
	UPPERCASE_K                                UtfChar = 75
	UPPERCASE_L                                UtfChar = 76
	UPPERCASE_M                                UtfChar = 77
	UPPERCASE_N                                UtfChar = 78
	UPPERCASE_O                                UtfChar = 79
	UPPERCASE_P                                UtfChar = 80
	UPPERCASE_Q                                UtfChar = 81
	UPPERCASE_R                                UtfChar = 82
	UPPERCASE_S                                UtfChar = 83
	UPPERCASE_T                                UtfChar = 84
	UPPERCASE_U                                UtfChar = 85
	UPPERCASE_V                                UtfChar = 86
	UPPERCASE_W                                UtfChar = 87
	UPPERCASE_X                                UtfChar = 88
	UPPERCASE_Y                                UtfChar = 89
	UPPERCASE_Z                                UtfChar = 90
	OPENING_SQUARE_BRACKET                     UtfChar = 91
	BACKSLASH                                  UtfChar = 92
	CLOSING_SQUARE_BRACKET                     UtfChar = 93
	CARET                                      UtfChar = 94
	UNDERSCORE                                 UtfChar = 95
	GRAVE_ACCENT                               UtfChar = 96
	LOWERCASE_A                                UtfChar = 97
	LOWERCASE_B                                UtfChar = 98
	LOWERCASE_C                                UtfChar = 99
	LOWERCASE_D                                UtfChar = 100
	LOWERCASE_E                                UtfChar = 101
	LOWERCASE_F                                UtfChar = 102
	LOWERCASE_G                                UtfChar = 103
	LOWERCASE_H                                UtfChar = 104
	LOWERCASE_I                                UtfChar = 105
	LOWERCASE_J                                UtfChar = 106
	LOWERCASE_K                                UtfChar = 107
	LOWERCASE_L                                UtfChar = 108
	LOWERCASE_M                                UtfChar = 109
	LOWERCASE_N                                UtfChar = 110
	LOWERCASE_O                                UtfChar = 111
	LOWERCASE_P                                UtfChar = 112
	LOWERCASE_Q                                UtfChar = 113
	LOWERCASE_R                                UtfChar = 114
	LOWERCASE_S                                UtfChar = 115
	LOWERCASE_T                                UtfChar = 116
	LOWERCASE_U                                UtfChar = 117
	LOWERCASE_V                                UtfChar = 118
	LOWERCASE_W                                UtfChar = 119
	LOWERCASE_X                                UtfChar = 120
	LOWERCASE_Y                                UtfChar = 121
	LOWERCASE_Z                                UtfChar = 122
	OPENING_CURLY_BRACKET                      UtfChar = 123
	VERTICAL_BAR                               UtfChar = 124
	CLOSING_CURLY_BRACKET                      UtfChar = 125
	TILDE                                      UtfChar = 126
	DELETE                                     UtfChar = 127
	EURO_SIGN                                  UtfChar = 128
	UNUSED_1                                   UtfChar = 129
	SINGLE_LOW_9_QUOTATION_MARK                UtfChar = 130
	LOWERCASE_F_WITH_HOOK                      UtfChar = 131
	DOUBLE_LOW_9_QUOTATION_MARK                UtfChar = 132
	HORIZONTAL_ELLIPSIS                        UtfChar = 133
	DAGGER                                     UtfChar = 134
	DOUBLE_DAGGER                              UtfChar = 135
	MODIFIER_LETTER_CIRCUMFLEX_ACCENT          UtfChar = 136
	PER_MILE_SIGN                              UtfChar = 137
	UPPERCASE_S_WITH_CARON                     UtfChar = 138
	SINGLE_LEFT_POINTING_ANGLE_QUOTATION       UtfChar = 139
	LATIN_CAPTIAL_LIGATURE_OE                  UtfChar = 140
	UNUSED_2                                   UtfChar = 141
	UPPERCASE_Z_WITH_CARON                     UtfChar = 142
	UNUSED_3                                   UtfChar = 143
	UNUSED_4                                   UtfChar = 144
	LEFT_SINGLE_QUOTATION_MARK                 UtfChar = 145
	RIGHT_SINGLE_QUOTATION_MARK                UtfChar = 146
	LEFT_DOUBLE_QUOTATION_MARK                 UtfChar = 147
	RIGHT_DOUBLE_QUOTATION_MARK                UtfChar = 148
	BULLET                                     UtfChar = 149
	EN_DASH                                    UtfChar = 150
	EM_DASH                                    UtfChar = 151
	SMALL_TILDE                                UtfChar = 152
	TRADE_MARK_SIGN                            UtfChar = 153
	LATIN_SMALL_S_WITH_CARON                   UtfChar = 154
	SINGLE_RIGHT_POINTING_ANBLE_QUOTATION_MARK UtfChar = 155
	LOWERCASE_LIGATURE_OE                      UtfChar = 156
	UNUSED_5                                   UtfChar = 157
	LOWERCASE_Z_WITH_CARON                     UtfChar = 158
	UPPERCASE_Y_WITH_DIAERESIS                 UtfChar = 159
	NON_BREAKING_SPACE                         UtfChar = 160
	INVERTED_EXCLAMATION_MARK                  UtfChar = 161
	CENT_SIGN                                  UtfChar = 162
	POUND_SIGN                                 UtfChar = 163
	CURRENCY_SIGN                              UtfChar = 164
	YEN_SIGN                                   UtfChar = 165
	BROKEN_VERTICAL_BAR                        UtfChar = 166
	SECTION_SIGN                               UtfChar = 167
	SPACING_DIARESIS_UMLAUT                    UtfChar = 168
	COPYRIGHT_SIGN                             UtfChar = 169
	FEMININE_ORDINAL_INDICATOR                 UtfChar = 170
	LEFT_DOUBLE_ANGLE_QUOTES                   UtfChar = 171
	NEGATION                                   UtfChar = 172
	SOFT_HYPHEN                                UtfChar = 173
	REGISTERED_TRADE_MARK_SIGN                 UtfChar = 174
	SPACING_MACRON_OVERLINE                    UtfChar = 175
	DEGREE_SIGN                                UtfChar = 176
	PLUS_OR_MINUS_SIGN                         UtfChar = 177
	SUPERSCRIPT_2_SQUARED                      UtfChar = 178
	SUPERSCRIPT_3_CUBED                        UtfChar = 179
	ACUTE_ACCENT_SPACING_ACCUTE                UtfChar = 180
	MICRO_SIGN                                 UtfChar = 181
	PILCROW_SIGN_PARAGRAPH_SIGN                UtfChar = 182
	MIDDLE_DOT_GEORGIAN_COMMA                  UtfChar = 183
	SPACING_CEDILLA                            UtfChar = 184
	SUPERSCRIPT_1                              UtfChar = 185
	MASCULINE_ORDINAL_INDICATOR                UtfChar = 186
	RIGHT_DOUBLE_ANGLE_QUOTES                  UtfChar = 187
	FRACTION_1_4                               UtfChar = 188
	FRACTION_1_2                               UtfChar = 189
	FRACTION_3_4                               UtfChar = 190
	INVERTED_QUESTION_MARK                     UtfChar = 191
	UPPERCASE_A_WITH_GRAVE                     UtfChar = 192
	UPPERCASE_A_WITH_ACUTE                     UtfChar = 193
	UPPERCASE_A_WITH_CIRCUMFLEX                UtfChar = 194
	UPPERCASE_A_WITH_TILDE                     UtfChar = 195
	UPPERCASE_A_WITH_DIARESIS                  UtfChar = 196
	UPPERCASE_A_WITH_RING_ABOVE                UtfChar = 197
	UPPERCASE_AE                               UtfChar = 198
	UPPERCASE_C_WITH_CEDILLA                   UtfChar = 199
	UPPERCASE_e_WITH_GRAVE                     UtfChar = 200
	UPPERCASE_E_WITH_ACUTE                     UtfChar = 201
	UPPERCASE_E_WITH_circumflex                UtfChar = 202
	UPPERCASE_E_WITH_DIAERESIS                 UtfChar = 203
	UPPERCASE_I_WITH_GRAVE                     UtfChar = 204
	UPPERCASE_I_WITH_ACUTE                     UtfChar = 205
	UPPERCASE_I_WITH_circumflex                UtfChar = 206
	UPPERCASE_I_WITH_DIAERESIS                 UtfChar = 207
	UPPERCASE_ETH                              UtfChar = 208
	UPPERCASE_N_WITH_TILDE                     UtfChar = 209
	UPPERCASE_O_WITH_GRAVE                     UtfChar = 210
	UPPERCASE_O_WITH_ACUTE                     UtfChar = 211
	UPPERCASE_O_WITH_circumflex                UtfChar = 212
	UPPERCASE_O_WITH_TILDE                     UtfChar = 213
	UPPERCASE_O_WITH_DIAERESIS                 UtfChar = 214
	MULTIPLICATION_SIGN                        UtfChar = 215
	UPPERCASE_O_WITH_SLASH                     UtfChar = 216
	UPPERCASE_U_WITH_GRAVE                     UtfChar = 217
	UPPERCASE_U_WITH_ACUTE                     UtfChar = 218
	UPPERCASE_U_WITH_circumflex                UtfChar = 219
	UPPERCASE_U_WITH_DIAERESIS                 UtfChar = 220
	UPPERCASE_Y_WITH_ACUTE                     UtfChar = 221
	UPPERCASE_THORN                            UtfChar = 222
	LOWERCASE_SHARP_S                          UtfChar = 223
	LOWERCASE_A_WITH_GRAVE                     UtfChar = 224
	LOWERCASE_A_WITH_ACUTE                     UtfChar = 225
	LOWERCASE_A_WITH_circumflex                UtfChar = 226
	LOWERCASE_A_WITH_TILDE                     UtfChar = 227
	LOWERCASE_A_WITH_DIAERESIS                 UtfChar = 228
	LOWERCASE_A_WITH_RING_ABOVE                UtfChar = 229
	LOWERCASE_AE                               UtfChar = 230
	LOWERCASE_C_WITH_CEDILLA                   UtfChar = 231
	LOWERCASE_E_WITH_GRAVE                     UtfChar = 232
	LOWERCASE_E_WITH_ACUTE                     UtfChar = 233
	LOWERCASE_E_WITH_circumflex                UtfChar = 234
	LOWERCASE_E_WITH_DIAERESIS                 UtfChar = 235
	LOWERCASE_I_WITH_GRAVE                     UtfChar = 236
	LOWERCASE_I_WITH_ACUTE                     UtfChar = 237
	LOWERCASE_I_WITH_circumflex                UtfChar = 238
	LOWERCASE_I_WITH_DIAERESIS                 UtfChar = 239
	LOWERCASE_ETH                              UtfChar = 240
	LOWERCASE_N_WITH_TILDE                     UtfChar = 241
	LOWERCASE_O_WITH_GRAVE                     UtfChar = 242
	LOWERCASE_O_WITH_ACUTE                     UtfChar = 243
	LOWERCASE_O_WITH_circumflex                UtfChar = 244
	LOWERCASE_O_WITH_TILDE                     UtfChar = 245
	LOWERCASE_O_WITH_DIAERESIS                 UtfChar = 246
	DIVISION_SIGN                              UtfChar = 247
	LOWERCASE_O_WITH_SLASH                     UtfChar = 248
	LOWERCASE_U_WITH_GRAVE                     UtfChar = 249
	LOWERCASE_U_WITH_ACUTE                     UtfChar = 250
	LOWERCASE_U_WITH_circumflex                UtfChar = 251
	LOWERCASE_U_WITH_DIAERESIS                 UtfChar = 252
	LOWERCASE_Y_WITH_ACUTE                     UtfChar = 253
	LOWERCASE_THORN                            UtfChar = 254
	LOWERCASE_Y_WITH_DIAERESIS                 UtfChar = 255
)
