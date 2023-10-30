package validator

type SliceValidator[T any] struct {
	fns []Validatable
}
