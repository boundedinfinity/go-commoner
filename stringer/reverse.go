package stringer

func Reverse[T ~string](s T) string {
	var rs []byte

	for i := len(s) - 1; i >= 0; i-- {
		rs = append(rs, s[i])
	}

	return string(rs)
}
