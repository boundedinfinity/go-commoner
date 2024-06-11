package pipeliner

func Build[T any]() *pipeline[T] {
	return &pipeline[T]{
		fns: make([]func(T) (T, error), 0),
	}
}

type pipeline[T any] struct {
	fns []func(T) (T, error)
}

func (p *pipeline[T]) Then(fn func(T) (T, error)) *pipeline[T] {
	p.fns = append(p.fns, fn)
	return p
}

func (p *pipeline[T]) MustThen(fn func(T) T) *pipeline[T] {
	return p.Then(func(elem T) (T, error) {
		return fn(elem), nil
	})
}

func (p pipeline[T]) Run(elem T) (T, error) {
	var err error

	for _, fn := range p.fns {
		elem, err = fn(elem)

		if err != nil {
			break
		}
	}

	return elem, err
}

func (p pipeline[T]) MustRun(elem T) T {
	elem, err := p.Run(elem)

	if err != nil {
		panic(err)
	}

	return elem
}
