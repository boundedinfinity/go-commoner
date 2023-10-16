package validater

type Validater interface {
	Validate(groups ...string) error
}

type validater[T any] struct {
	fns  []func(T) error
	item T
}

func (t *validater[T]) Validate(groups ...string) error {
	for _, fn := range t.fns {
		if err := fn(t.item); err != nil {
			return err
		}
	}

	return nil
}
