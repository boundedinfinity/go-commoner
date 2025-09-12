package slicer

import (
	"fmt"
	gstring "strings"
)

func Join[T any](sep string, elems ...T) string {
	fn := func(elem T) string { return fmt.Sprintf("%v", elem) }
	return JoinFn(fn, sep, elems...)
}

func JoinFn[T any](fn func(T) string, sep string, elems ...T) string {
	return JoinIFn(wrapIndexFunc(fn), sep, elems...)
}

func JoinIFn[T any](fn func(int, T) string, sep string, elems ...T) string {
	fn2 := func(i int, elem T) (string, error) { return fn(i, elem), nil }
	return ignoreErr(JoinIErrFn(fn2, sep, elems...))
}

func JoinErrFn[T any](fn func(T) (string, error), sep string, elems ...T) (string, error) {
	return JoinIErrFn(wrapIndexAndErr(fn), sep, elems...)
}

func JoinIErrFn[T any](fn func(int, T) (string, error), sep string, elems ...T) (string, error) {
	result, err := MapErrI(fn, elems...)

	if err != nil {
		return "", err
	}

	return gstring.Join(result, sep), nil
}
