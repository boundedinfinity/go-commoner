package slicer

func New[T any](items ...T) Slicer[T] {
	return Slicer[T](items)
}

type Slicer[T any] []T
