package chain

import "github.com/boundedinfinity/go-commoner/stringer"

func StringRemover[T ~string](remove string) ChainFn[T] {
	return func(s T) T {
		res := stringer.Replace[T, string, string](s, remove, "")
		return T(res)
	}
}
