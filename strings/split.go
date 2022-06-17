package strings

func Split[T ~string](s T, sep string) []string {
	return SplitAny(s, sep)
}

func SplitAny[T ~string](s T, cs ...string) []string {
	return SplitFn(s, func(x string) bool {
		for _, c := range cs {
			if c == string(x) {
				return true
			}
		}

		return false
	})
}

func SplitFn[T ~string](s T, fn func(string) bool) []string {
	var o []string
	start := 0

	for i, x := range s {
		if fn(string(x)) {
			o = append(o, string(s[start:i]))
			start = i + 1
		}
	}

	o = append(o, string(s[start:]))

	return o
}
