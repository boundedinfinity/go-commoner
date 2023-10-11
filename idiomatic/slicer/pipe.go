package slicer

type Pipe[T any] struct {
	fns []func(T) T
}

func NewPipe[T any]() *Pipe[T] {
	return &Pipe[T]{
		fns: make([]func(T) T, 0),
	}
}

func (p *Pipe[T]) Then(fn func(T) T) *Pipe[T] {
	p.fns = append(p.fns, fn)
	return p
}

func (p *Pipe[T]) Run(items ...T) []T {
	output := append([]T{}, items...)

	for _, fn := range p.fns {
		output = Map(output, fn)
	}

	return output
}

func (p *Pipe[T]) RunSingle(item T) T {
	items := p.Run(item)

	if len(items) > 0 {
		return items[0]
	}

	var zero T
	return zero
}

type PipeErr[T any] struct {
	fns []func(T) (T, error)
}

func NewPipeErr[T any]() *PipeErr[T] {
	return &PipeErr[T]{
		fns: make([]func(T) (T, error), 0),
	}
}

func (p *PipeErr[T]) Then(fn func(T) (T, error)) *PipeErr[T] {
	p.fns = append(p.fns, fn)
	return p
}

func (p *PipeErr[T]) Run(items ...T) ([]T, error) {
	output := append([]T{}, items...)
	var err error

	for _, fn := range p.fns {
		output, err = MapErr(output, fn)

		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func (p *PipeErr[T]) RunSingle(item T) (T, error) {
	items, err := p.Run(item)

	if len(items) > 0 {
		return items[0], err
	}

	var zero T
	return zero, err
}
