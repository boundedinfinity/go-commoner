package stringer

func TruncateEndEllipse[T ~string](s T, w int) string {
	return TruncateEnd(s, w, "...")
}

func TruncateEnd[T ~string](s T, w int, m string) string {
	length := len(s)

	if length <= w {
		return string(s)
	}

	ml := len(m)

	if length+ml <= w {
		return string(s)
	}

	sub := string(s[:w-ml]) + m

	return sub
}

func TruncateBeginEllipse[T ~string](s T, w int) string {
	return TruncateBegin(s, w, "...")
}

func TruncateBegin[T ~string](s T, w int, m string) string {
	length := len(s)

	if length <= w {
		return string(s)
	}

	ml := len(m)

	if length+ml <= w {
		return string(s)
	}

	sub := m + string(s[length-w+ml:])

	return sub
}
