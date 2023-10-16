package validater

type SliceValidator[T any] struct {
	fns []Validater
}
