package strings

func Abbreviate[T ~string](s T, w int) string {
	return AbbreviateWith(s, w, "...")
}

func AbbreviateWith[T ~string](s T, w int, m string) string {
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

func AbbreviateLeft[T ~string](s T, w int) string {
	return AbbreviateLeftWith(s, w, "...")
}

func AbbreviateLeftWith[T ~string](s T, w int, m string) string {
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
