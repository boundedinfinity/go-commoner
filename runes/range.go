package runes

func Range(s, e rune) []rune {
	var bs []rune

	for i := s; i <= e; i++ {
		bs = append(bs, i)
	}

	return bs
}
