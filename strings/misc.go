package strings

func FirstNotEmpty[T ~string](ss ...T) T {
	for _, s := range ss {
		if !IsEmpty(s) {
			return s
		}
	}

	return ""
}

func OrElse[T ~string](s, d T) T {
	if s == "" {
		return d
	}

	return s
}

func DefaultIfEmpty[T ~string](s, d T) T {
	return OrElse(s, d)
}

func IsEmpty[T ~string](s T) bool {
	if s == "" {
		return true
	}

	return false
}

func IsBlank[T ~string](s T) bool {
	return IsEmpty(s)
}

func Abbreviate[T ~string](s T, w int) T {
	return AbbreviateWith(s, w, "...")
}

func AbbreviateWith[T ~string](s T, w int, m string) T {
	sl := len(s)

	if sl <= w {
		return s
	}

	ml := len(m)

	if sl+ml <= w {
		return s
	}

	sub := s[:w-ml] + T(m)

	return sub
}
