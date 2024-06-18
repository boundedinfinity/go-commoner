package optioner

func Run[T any](opt Option[T], fn func(T) Option[T]) Option[T] {
	if opt.Defined() {
		return fn(opt.Get())
	}

	return None[T]()
}

func (t Option[T]) Run(fn func(T) Option[T]) Option[T] {
	return Run(t, fn)
}
