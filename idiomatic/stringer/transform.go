package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/errorer"

func TarnsformFunc[T ~string](s T, fns ...func(string) string) string {
	native := string(s)

	for _, fn := range fns {
		if fn != nil {
			native = fn(native)
		}
	}

	return native
}

func TransformAllFunc[T ~string](items []T, fns ...func(string) string) []string {
	results := make([]string, len(items))

	for i := range items {
		results[i] = TarnsformFunc(items[i], fns...)
	}

	return results
}

var (
	ErrTransform   = errorer.New("transform error")
	errTransformFn = errorer.Func(ErrStringer, ErrTransform)
)

func TransformFuncErr[T ~string](s T, fns ...func(string) (string, error)) (string, error) {
	native := string(s)
	var err error

	for i, fn := range fns {
		if fn != nil {
			if native, err = fn(native); err != nil {
				err = errTransformFn("fns[%d] : %w", i, err)
				break
			}
		}
	}

	return native, err
}

func TransformAllFuncErr[T ~string](items []T, fns ...func(string) (string, error)) ([]string, error) {
	results := make([]string, len(items))
	var err error

	for i := range items {
		if err != nil {
			break
		}

		native := string(items[i])

		for j, fn := range fns {
			if fn != nil {
				if native, err = fn(native); err != nil {
					err = errTransformFn("item[%d] : fns[%d] : %w", i, j, err)
					break
				}
			}
		}
	}

	return results, err
}
