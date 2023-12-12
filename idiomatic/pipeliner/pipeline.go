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
	return p.Then(func(item T) (T, error) {
		return fn(item), nil
	})
}

func (p pipeline[T]) Run(item T) (T, error) {
	var err error

	for _, fn := range p.fns {
		item, err = fn(item)

		if err != nil {
			break
		}
	}

	return item, err
}

func (p pipeline[T]) MustRun(item T) T {
	item, err := p.Run(item)

	if err != nil {
		panic(err)
	}

	return item
}
