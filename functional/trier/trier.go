package trier

type Try[T any] struct {
	Result T     `json:"result"`
	Err    error `json:"error"`
}

func (t Try[T]) Failed() bool {
	return t.Err != nil
}

func (t Try[T]) Succeeded() bool {
	return !t.Failed()
}

func (t Try[T]) Error() string {
	return t.Err.Error()
}
