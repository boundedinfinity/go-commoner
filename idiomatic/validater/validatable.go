package validator

type Validatable interface {
	Validate(groups ...string) error
	Groups() []string
}

type validatable[T any] struct {
	fns    []func(T) error
	groups []string
	elem   T
}

func (t *validatable[T]) Groups() []string {
	return t.groups
}

func (t *validatable[T]) Validate(groups ...string) error {
	for _, fn := range t.fns {
		if err := fn(t.elem); err != nil {
			return err
		}
	}

	return nil
}
