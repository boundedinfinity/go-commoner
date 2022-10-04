package chain

import "github.com/boundedinfinity/go-commoner/stringer"

func StringRemover[T ~string](remove string) ChainFn[T] {
	return func(s T) T {
		res := stringer.Replace(s, remove, "")
		return T(res)
	}
}
