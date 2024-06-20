package mapper

func Upsert[T any, C comparable](m map[C]T, k C, t T) bool {
	fn := func(t T) C { return k }
	return UpsertFn(m, fn, t)
}

func UpsertFn[T any, C comparable](m map[C]T, fn func(T) C, t T) bool {
	var ok bool

	if m != nil && fn != nil {
		k := fn(t)

		if _, ok = m[k]; !ok {
			m[k] = t
		}
	}

	return ok
}
