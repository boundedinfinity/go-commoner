package stringer

func Remove[T ~string](s T, remove string) string {
	return Replace[T](s, "", remove)
}
