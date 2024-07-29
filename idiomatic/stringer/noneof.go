package stringer

func NoneOf[T ~string, S ~string](s S, elems ...T) bool {
	return !AnyOf(s, elems...)
}

func NoneOfIgnoreCase[T ~string, S ~string](s S, elems ...T) bool {
	return !AnyOfIgnoreCase(s, elems...)
}
