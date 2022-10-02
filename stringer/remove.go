package stringer

func Remove[T ~string](s T, remove string) string {
	return Replace(s, remove, "")
}
