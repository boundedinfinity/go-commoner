package strings

func Split[T ~string](s T) []T {
	return SplitAny[T](s, ' ')
}

func SplitAny[T ~string](s T, cs ...rune) []T {
	return SplitFn[T](s, func(x rune) bool {
		for _, c := range cs {
			if c == x {
				return true
			}
		}

		return false
	})
}

func SplitFn[T ~string](s T, fn func(rune) bool) []T {
	var o []T
	start := 0
	
	for i, x := range s {
		if fn(x) {
			o = append(o, s[start:i])
			start = i + 1
		}
	}

	o = append(o, s[start:])

	return o
}


func Join[T ~string](s []T, j T) T {
	var o T

	for i, x := range s {
		o += x

		if i != len(s) {
			o += j
		}
	}

	return o
}