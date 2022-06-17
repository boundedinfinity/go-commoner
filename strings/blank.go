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
