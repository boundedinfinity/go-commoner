package stringer

func TruncateEnd[T ~string](s T, w int) string {
	return TruncateEndWith(s, w, "...")
}

func TruncateEndWith[T ~string](s T, w int, m string) string {
	sl := len(s)

	if sl <= w {
		return string(s)
	}

	ml := len(m)

	if sl+ml <= w {
		return string(s)
	}

	sub := string(s[:w-ml]) + m

	return sub
}

func TruncateBegin[T ~string](s T, w int) string {
	return TruncateBeginWith(s, w, "...")
}

func TruncateBeginWith[T ~string](s T, w int, m string) string {
	sl := len(s)

	if sl <= w {
		return string(s)
	}

	ml := len(m)

	if sl+ml <= w {
		return string(s)
	}

	sub := m + string(s[sl-w+ml:])

	return sub
}
