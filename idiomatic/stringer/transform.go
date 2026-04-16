package stringer

func TarnsformFunc[T ~string](s T, fns ...func(string) string) string {
	native := string(s)

	for _, fn := range fns {
		if fn != nil {
			native = fn(native)
		}
	}

	return native
}

func TransformAllFunc[T ~string](items []T, fns ...func(string) string) []string {
	results := make([]string, len(items))

	for i := range items {
		results[i] = TarnsformFunc(items[i], fns...)
	}

	return results
}
