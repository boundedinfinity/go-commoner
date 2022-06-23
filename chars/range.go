package chars

func Range(s, e byte) []byte {
	var bs []byte

	for i := s; i <= e; i++ {
		bs = append(bs, i)
	}

	return bs
}
