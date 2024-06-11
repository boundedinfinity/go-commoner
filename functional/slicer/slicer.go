package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

type Slicer[T any] optioner.Option[[]T]

// func (t Slicer[T]) Len() int {
// 	return len(t.Get())
// }

// func (t *Slicer[T]) Append(elem T) {
// 	t.Option = optioner.OfSlice(append(t.Get(), elem))
// }
