package chain

import (
	"github.com/boundedinfinity/go-commoner/chain"
	"github.com/boundedinfinity/go-commoner/trier"
)

func New[T any]() *Chain[T] {
	return &Chain[T]{
		internal: chain.NewError[T](),
	}
}

type Chain[T any] struct {
	internal *chain.ChainErr[T]
}

func (t *Chain[T]) Append(fn chain.ChainErrFn[T]) *Chain[T] {
	t.internal.Append(fn)
	return t
}

func (t *Chain[T]) AppendNoError(fn chain.ChainFn[T]) *Chain[T] {
	t.internal.AppendNoErr(fn)
	return t
}

func (t *Chain[T]) Prepend(fn chain.ChainErrFn[T]) *Chain[T] {
	t.internal.Prepend(fn)
	return t
}

func (t *Chain[T]) RunList(items []T) trier.Try[[]T] {
	res, err := t.internal.RunList(items)
	return trier.Complete(res, err)
}

func (t *Chain[T]) RunSingle(item T) trier.Try[T] {
	res := t.RunList([]T{item})
	return trier.Complete(res.Result[0], res.Error)
}
