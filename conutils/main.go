package conutils

type WorkResult[T any] struct {
	data T
	err error
}

type WorkFn[T any]() error

func Routine[T any](fn WorkFn[T]) chan WorkResult[T] {
	ch := <- WorkResult[T]
	var wg sync.WaitGroup 

	go func() {
		defer wg.Done()
		defer close(ch)

		data, err := fn()

		ch <- WorkResult {
			data: data,
			err: err,
		}
	}

	return ch
}