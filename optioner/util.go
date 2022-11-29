package optioner

func FirstOf[T any](options ...Option[T]) Option[T] {
	for _, option := range options {
		if option.Defined() {
			return option
		}
	}

	return None[T]()
}

func LastOf[T any](options ...Option[T]) Option[T] {
	for i := len(options) - 1; i >= 0; i-- {
		if options[i].Defined() {
			return options[i]
		}
	}

	return None[T]()
}
