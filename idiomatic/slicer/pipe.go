package slicer

type Pipe[T any] struct {
	fns []func(int, T) T
}

func NewPipe[T any]() *Pipe[T] {
	return &Pipe[T]{
		fns: make([]func(int, T) T, 0),
	}
}

func (p *Pipe[T]) Then(fn func(int, T) T) *Pipe[T] {
	p.fns = append(p.fns, fn)
	return p
}

func (p *Pipe[T]) List(elems ...T) []T {
	output := append([]T{}, elems...)

	for _, fn := range p.fns {
		output = Map(fn, output...)
	}

	return output
}

func (p *Pipe[T]) Single(elem T) T {
	elems := p.List(elem)

	if len(elems) > 0 {
		return elems[0]
	}

	var zero T
	return zero
}

type PipeErr[T any] struct {
	fns []func(int, T) (T, error)
}

func NewPipeErr[T any]() *PipeErr[T] {
	return &PipeErr[T]{
		fns: make([]func(int, T) (T, error), 0),
	}
}

func (p *PipeErr[T]) Then(fn func(int, T) (T, error)) *PipeErr[T] {
	p.fns = append(p.fns, fn)
	return p
}

func (p *PipeErr[T]) List(elems ...T) ([]T, error) {
	output := append([]T{}, elems...)
	var err error

	for _, fn := range p.fns {
		output, err = MapErr(fn, output...)

		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func (p *PipeErr[T]) Single(elem T) (T, error) {
	elems, err := p.List(elem)

	if len(elems) > 0 {
		return elems[0], err
	}

	var zero T
	return zero, err
}
