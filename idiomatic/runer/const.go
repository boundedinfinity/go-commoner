package runer

const (
	NULL                      rune = 0x00
	START_OF_HEADER           rune = 0x01
	START_OF_TEXT             rune = 0x02
	END_OF_TEXT               rune = 0x03
	END_OF_TRANSMISSION       rune = 0x04
	ENQUIRY                   rune = 0x05
	ACKNOWLEDGE               rune = 0x06
	BELL                      rune = 0x07
	BACKSPACE                 rune = 0x08
	HORIZONTAL_TAB            rune = 0x09
	LINE_FEED                 rune = 0x0A
	VERTICAL_TAB              rune = 0x0B
	FORM_FEED                 rune = 0x0C
	CARRIAGE_RETURN           rune = 0x0D
	SHIFT_OUT                 rune = 0x0E
	SHIFT_IN                  rune = 0x0F
	DATA_LINK_ESCAPE          rune = 0x10
	DEVICE_CONTROL_1          rune = 0x11
	DEVICE_CONTROL_2          rune = 0x12
	DEVICE_CONTROL_3          rune = 0x13
	DEVICE_CONTROL_4          rune = 0x14
	NEGATIVE_ACKNOLEDGE       rune = 0x15
	SYNCHRONIZE               rune = 0x16
	END_OF_TRANSMISSION_BLOCK rune = 0x17
	CANCEL                    rune = 0x18
	END_OF_MEDIUM             rune = 0x19
	SUBSTITUE                 rune = 0x1A
	ESCAPE                    rune = 0x1B
	FILE_SEPARATOR            rune = 0x1C
	GROUP_SEPARATOR           rune = 0x1D
	RECORD_SEPARATOR          rune = 0x1E
	UNIT_SEPARATOR            rune = 0x1F
	SPACE                     rune = 0x20
	EXCLAMATION_MARK          rune = 0x21
	DOUBLE_QUOTE              rune = 0x22
	NUMBER                    rune = 0x23
	DOLLAR                    rune = 0x24
	PERCENT                   rune = 0x25
	AMPERSAND                 rune = 0x26
	SINGLE_QUOTE              rune = 0x27
	LEFT_PARENTHESIS          rune = 0x28
	RIGHT_PARENTHESIS         rune = 0x29
	ASTERISK                  rune = 0x2A
	PLUS                      rune = 0x2B
	COMMA                     rune = 0x2C
	MINUS                     rune = 0x2D
	PERIOD                    rune = 0x2E
	SLASH                     rune = 0x2F
	ZERO                      rune = 0x30
	ONE                       rune = 0x31
	TWO                       rune = 0x32
	THREE                     rune = 0x33
	FOUR                      rune = 0x34
	FIVE                      rune = 0x35
	SIX                       rune = 0x36
	SEVEN                     rune = 0x37
	EIGHT                     rune = 0x38
	NINE                      rune = 0x39
	COLON                     rune = 0x3A
	SEMICOLON                 rune = 0x3B
	LESS_THAN                 rune = 0x3C
	EQUAL_SIGN                rune = 0x3D
	GREATER_THAN              rune = 0x3E
	QUESTION_MARK             rune = 0x3F
	AT_SIGN                   rune = 0x40
	UPPERCASE_A               rune = 0x41
	UPPERCASE_B               rune = 0x42
	UPPERCASE_C               rune = 0x43
	UPPERCASE_D               rune = 0x44
	UPPERCASE_E               rune = 0x45
	UPPERCASE_F               rune = 0x46
	UPPERCASE_G               rune = 0x47
	UPPERCASE_H               rune = 0x48
	UPPERCASE_I               rune = 0x49
	UPPERCASE_J               rune = 0x4A
	UPPERCASE_K               rune = 0x4B
	UPPERCASE_L               rune = 0x4C
	UPPERCASE_M               rune = 0x4D
	UPPERCASE_N               rune = 0x4E
	UPPERCASE_O               rune = 0x4F
	UPPERCASE_P               rune = 0x50
	UPPERCASE_Q               rune = 0x51
	UPPERCASE_R               rune = 0x52
	UPPERCASE_S               rune = 0x53
	UPPERCASE_T               rune = 0x54
	UPPERCASE_U               rune = 0x55
	UPPERCASE_V               rune = 0x56
	UPPERCASE_W               rune = 0x57
	UPPERCASE_X               rune = 0x58
	UPPERCASE_Y               rune = 0x59
	UPPERCASE_Z               rune = 0x5A
	LEFT_SQUARE_BRACKET       rune = 0x5B
	BACKSLASH                 rune = 0x5C
	RIGHT_SQUARE_BRACKET      rune = 0x5D
	CARET                     rune = 0x5E
	UNDERSCORE                rune = 0x5F
	GRAVE                     rune = 0x60
	LOWERCASE_A               rune = 0x61
	LOWERCASE_B               rune = 0x62
	LOWERCASE_C               rune = 0x63
	LOWERCASE_D               rune = 0x64
	LOWERCASE_E               rune = 0x65
	LOWERCASE_F               rune = 0x66
	LOWERCASE_G               rune = 0x67
	LOWERCASE_H               rune = 0x68
	LOWERCASE_I               rune = 0x69
	LOWERCASE_J               rune = 0x6A
	LOWERCASE_K               rune = 0x6B
	LOWERCASE_L               rune = 0x6C
	LOWERCASE_M               rune = 0x6D
	LOWERCASE_N               rune = 0x6E
	LOWERCASE_O               rune = 0x6F
	LOWERCASE_p               rune = 0x70
	LOWERCASE_Q               rune = 0x71
	LOWERCASE_R               rune = 0x72
	LOWERCASE_S               rune = 0x73
	LOWERCASE_T               rune = 0x74
	LOWERCASE_U               rune = 0x75
	LOWERCASE_V               rune = 0x76
	LOWERCASE_W               rune = 0x77
	LOWERCASE_X               rune = 0x78
	LOWERCASE_Y               rune = 0x79
	LOWERCASE_Z               rune = 0x7A
	LEFT_CURLY_BRACKET        rune = 0x7B
	VERTICAL_BAR              rune = 0x7C
	RIGHT_CURLY_BRACKET       rune = 0x7D
	TILDE                     rune = 0x7E
	DELETE                    rune = 0x7F
)