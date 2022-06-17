package bytes

func IsInteger(v byte) bool {
	return '0' <= v && v <= '9'
}
