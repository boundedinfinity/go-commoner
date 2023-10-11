package trier

type Try[T any] struct {
	Result T     `json:"result"`
	Error  error `json:"error"`
}

func (t Try[T]) Failure() bool {
	return t.Error != nil
}

func (t Try[T]) Success() bool {
	return !t.Failure()
}
