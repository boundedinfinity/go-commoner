package slices

func Times[V any, O any](v V, s int, fn func(V) O) []O {
	var os []O

	for i := 0; i < s; i++ {
		os = append(os, fn(v))
	}

	return os
}
