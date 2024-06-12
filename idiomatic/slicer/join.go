package slicer

import (
	"fmt"
	gstring "strings"
)

func Join[T any](sep string, elems ...T) string {
	fn := func(_ int, elem T) string { return fmt.Sprintf("%v", elem) }
	return JoinFn(fn, sep, elems...)
}

func JoinFn[T any](fn func(int, T) string, sep string, elems ...T) string {
	fn2 := func(i int, elem T) (string, error) {
		return fn(i, elem), nil
	}

	result, _ := JoinErrFn(fn2, sep, elems...)
	return result
}

func JoinErrFn[T any](fn func(int, T) (string, error), sep string, elems ...T) (string, error) {
	result, err := MapErr(fn, elems...)

	if err != nil {
		return "", err
	}

	return gstring.Join(result, sep), nil
}
