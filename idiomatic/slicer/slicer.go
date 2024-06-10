package slicer

type Slicer[T any] []T

func (t Slicer[T]) Len() int {
	return len(t)
}

func (t *Slicer[T]) Append(items ...T) {
	*t = append(*t, items...)
}

func (t Slicer[T]) Slice() []T {
	return t
}

// ////////////////////////////////////////////////////////////////////////////
// Chunk
// ////////////////////////////////////////////////////////////////////////////

func (t Slicer[T]) Chunk(size int, items ...T) [][]T {
	return Chunk(size, t...)
}

// ////////////////////////////////////////////////////////////////////////////
// Fold
// ////////////////////////////////////////////////////////////////////////////

func (t Slicer[T]) FoldLeft(initial T, fn func(T, T) T, items ...T) T {
	return FoldLeft(initial, fn, t...)
}

func (t Slicer[T]) FoldLeftI(initial T, fn func(int, T, T) T, items ...T) T {
	return FoldLeftI(initial, fn, t...)
}

func (t Slicer[T]) FoldLeftErr(initial T, fn func(T, T) (T, error), items ...T) (T, error) {
	return FoldLeftErr(initial, fn, t...)
}

func (t Slicer[T]) FoldLeftErrI(initial T, fn func(int, T, T) (T, error), items ...T) (T, error) {
	return FoldLeftErrI(initial, fn, t...)
}

func (t Slicer[T]) FoldRight(initial T, fn func(T, T) T, items ...T) T {
	return FoldRight(initial, fn, t...)
}

func (t Slicer[T]) FoldRightI(initial T, fn func(int, T, T) T, items ...T) T {
	return FoldRightI(initial, fn, t...)
}

func (t Slicer[T]) FoldRightErr(initial T, fn func(T, T) (T, error), items ...T) (T, error) {
	return FoldRightErr(initial, fn, t...)
}

func (t Slicer[T]) FoldRightErrI(initial T, fn func(int, T, T) (T, error), items ...T) (T, error) {
	return FoldRightErrI(initial, fn, t...)
}

// ////////////////////////////////////////////////////////////////////////////
// Map
// ////////////////////////////////////////////////////////////////////////////

func (t Slicer[T]) Map(fn func(T) T) []T {
	return Map(fn, t...)
}

func (t Slicer[T]) MapI(fn func(int, T) T) []T {
	return MapI(fn, t...)
}

func (t Slicer[T]) MapErr(fn func(T) (T, error)) ([]T, error) {
	return MapErr(fn, t...)
}

func (t Slicer[T]) MapErrI(fn func(int, T) (T, error)) ([]T, error) {
	return MapErrI(fn, t...)
}
