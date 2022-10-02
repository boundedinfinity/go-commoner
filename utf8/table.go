package utf8

//go:generate enumer -path=./table.go

type Utf8Char byte

const (
	NULL                      Utf8Char = 0x00
	START_OF_HEADER           Utf8Char = 0x01
	START_OF_TEXT             Utf8Char = 0x02
	END_OF_TEXT               Utf8Char = 0x03
	END_OF_TRANSMISSION       Utf8Char = 0x04
	ENQUIRY                   Utf8Char = 0x05
	ACKNOWLEDGE               Utf8Char = 0x06
	BELL                      Utf8Char = 0x07
	BACKSPACE                 Utf8Char = 0x08
	HORIZONTAL_TAB            Utf8Char = 0x09
	LINE_FEED                 Utf8Char = 0x0A
	VERTICAL_TAB              Utf8Char = 0x0B
	FORM_FEED                 Utf8Char = 0x0C
	CARRIAGE_RETURN           Utf8Char = 0x0D
	SHIFT_OUT                 Utf8Char = 0x0E
	SHIFT_IN                  Utf8Char = 0x0F
	DATA_LINK_ESCAPE          Utf8Char = 0x10
	DEVICE_CONTROL_1          Utf8Char = 0x11
	DEVICE_CONTROL_2          Utf8Char = 0x12
	DEVICE_CONTROL_3          Utf8Char = 0x13
	DEVICE_CONTROL_4          Utf8Char = 0x14
	NEGATIVE_ACKNOLEDGE       Utf8Char = 0x15
	SYNCHRONIZE               Utf8Char = 0x16
	END_OF_TRANSMISSION_BLOCK Utf8Char = 0x17
	CANCEL                    Utf8Char = 0x18
	END_OF_MEDIUM             Utf8Char = 0x19
	SUBSTITUE                 Utf8Char = 0x1A
	ESCAPE                    Utf8Char = 0x1B
	FILE_SEPARATOR            Utf8Char = 0x1C
	GROUP_SEPARATOR           Utf8Char = 0x1D
	RECORD_SEPARATOR          Utf8Char = 0x1E
	UNIT_SEPARATOR            Utf8Char = 0x1F
	SPACE                     Utf8Char = 0x20
	EXCLAMATION_MARK          Utf8Char = 0x21
	DOUBLE_QUOTE              Utf8Char = 0x22
	NUMBER                    Utf8Char = 0x23
	DOLLAR                    Utf8Char = 0x24
	PERCENT                   Utf8Char = 0x25
	AMPERSAND                 Utf8Char = 0x26
	SINGLE_QUOTE              Utf8Char = 0x27
	LEFT_PARENTHESIS          Utf8Char = 0x28
	RIGHT_PARENTHESIS         Utf8Char = 0x29
	ASTERISK                  Utf8Char = 0x2A
	PLUS                      Utf8Char = 0x2B
	COMMA                     Utf8Char = 0x2C
	MINUS                     Utf8Char = 0x2D
	PERIOD                    Utf8Char = 0x2E
	SLASH                     Utf8Char = 0x2F
	NUMBER_ZERO               Utf8Char = 0x30
	NUMBER_ONE                Utf8Char = 0x31
	NUMBER_TWO                Utf8Char = 0x32
	NUMBER_THREE              Utf8Char = 0x33
	NUMBER_FOUR               Utf8Char = 0x34
	NUMBER_FIVE               Utf8Char = 0x35
	NUMBER_SIX                Utf8Char = 0x36
	NUMBER_SEVEN              Utf8Char = 0x37
	NUMBER_EIGHT              Utf8Char = 0x38
	NUMBER_NINE               Utf8Char = 0x39
	COLON                     Utf8Char = 0x3A
	SEMICOLON                 Utf8Char = 0x3B
	LESS_THAN                 Utf8Char = 0x3C
	EQUAL_SIGN                Utf8Char = 0x3D
	GREATER_THAN              Utf8Char = 0x3E
	QUESTION_MARK             Utf8Char = 0x3F
	AT_SIGN                   Utf8Char = 0x40
	UPPERCASE_A               Utf8Char = 0x41
	UPPERCASE_B               Utf8Char = 0x42
	UPPERCASE_C               Utf8Char = 0x43
	UPPERCASE_D               Utf8Char = 0x44
	UPPERCASE_E               Utf8Char = 0x45
	UPPERCASE_F               Utf8Char = 0x46
	UPPERCASE_G               Utf8Char = 0x47
	UPPERCASE_H               Utf8Char = 0x48
	UPPERCASE_I               Utf8Char = 0x49
	UPPERCASE_J               Utf8Char = 0x4A
	UPPERCASE_K               Utf8Char = 0x4B
	UPPERCASE_L               Utf8Char = 0x4C
	UPPERCASE_M               Utf8Char = 0x4D
	UPPERCASE_N               Utf8Char = 0x4E
	UPPERCASE_O               Utf8Char = 0x4F
	UPPERCASE_P               Utf8Char = 0x50
	UPPERCASE_Q               Utf8Char = 0x51
	UPPERCASE_R               Utf8Char = 0x52
	UPPERCASE_S               Utf8Char = 0x53
	UPPERCASE_T               Utf8Char = 0x54
	UPPERCASE_U               Utf8Char = 0x55
	UPPERCASE_V               Utf8Char = 0x56
	UPPERCASE_W               Utf8Char = 0x57
	UPPERCASE_X               Utf8Char = 0x58
	UPPERCASE_Y               Utf8Char = 0x59
	UPPERCASE_Z               Utf8Char = 0x5A
	LEFT_SQUARE_BRACKET       Utf8Char = 0x5B
	BACKSLASH                 Utf8Char = 0x5C
	RIGHT_SQUARE_BRACKET      Utf8Char = 0x5D
	CARET                     Utf8Char = 0x5E
	UNDERSCORE                Utf8Char = 0x5F
	GRAVE                     Utf8Char = 0x60
	LOWERCASE_A               Utf8Char = 0x61
	LOWERCASE_B               Utf8Char = 0x62
	LOWERCASE_C               Utf8Char = 0x63
	LOWERCASE_D               Utf8Char = 0x64
	LOWERCASE_E               Utf8Char = 0x65
	LOWERCASE_F               Utf8Char = 0x66
	LOWERCASE_G               Utf8Char = 0x67
	LOWERCASE_H               Utf8Char = 0x68
	LOWERCASE_I               Utf8Char = 0x69
	LOWERCASE_J               Utf8Char = 0x6A
	LOWERCASE_K               Utf8Char = 0x6B
	LOWERCASE_L               Utf8Char = 0x6C
	LOWERCASE_M               Utf8Char = 0x6D
	LOWERCASE_N               Utf8Char = 0x6E
	LOWERCASE_O               Utf8Char = 0x6F
	LOWERCASE_P               Utf8Char = 0x70
	LOWERCASE_Q               Utf8Char = 0x71
	LOWERCASE_R               Utf8Char = 0x72
	LOWERCASE_S               Utf8Char = 0x73
	LOWERCASE_T               Utf8Char = 0x74
	LOWERCASE_U               Utf8Char = 0x75
	LOWERCASE_V               Utf8Char = 0x76
	LOWERCASE_W               Utf8Char = 0x77
	LOWERCASE_X               Utf8Char = 0x78
	LOWERCASE_Y               Utf8Char = 0x79
	LOWERCASE_Z               Utf8Char = 0x7A
	LEFT_CURLY_BRACKET        Utf8Char = 0x7B
	VERTICAL_BAR              Utf8Char = 0x7C
	RIGHT_CURLY_BRACKET       Utf8Char = 0x7D
	TILDE                     Utf8Char = 0x7E
	DELETE                    Utf8Char = 0x7F
)
