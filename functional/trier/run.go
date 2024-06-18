package trier

func Run[T any](try Try[T], fn func(T) Try[T]) Try[T] {
	if try.Succeeded() {
		return fn(try.Value)
	}

	return try
}

func (t Try[T]) Run(fn func(T) Try[T]) Try[T] {
	return Run(t, fn)
}
