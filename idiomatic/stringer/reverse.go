package stringer

func Reverse[T ~string](s T) string {
	r1 := []rune(s)
	l := len(r1) - 1
	r2 := make([]rune, l+1)

	for i := range r1 {
		r2[i] = r1[l-i]
	}

	return string(r2)
}
