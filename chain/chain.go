package chain

type ChainFn[T any] func(T) T

func New[T any]() *Chain[T] {
	return &Chain[T]{
		chainWithErr: NewError[T](),
	}
}

type Chain[T any] struct {
	chainWithErr *ChainErr[T]
}

func (t *Chain[T]) Append(fn ChainFn[T]) *Chain[T] {
	t.chainWithErr.AppendNoErr(fn)
	return t
}

func (t *Chain[T]) Prepend(fn ChainFn[T]) *Chain[T] {
	t.chainWithErr.PrependNoErr(fn)
	return t
}

func (t *Chain[T]) RunSingle(item T) T {
	return t.RunList([]T{item})[0]
}

func (t *Chain[T]) RunList(items []T) []T {
	results, _ := t.chainWithErr.RunList(items)
	return results
}
