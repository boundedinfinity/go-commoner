package slicer

import (
	"fmt"
	gstring "strings"
)

func Join[T any](sep string, elems ...T) string {
	return JoinFn(sep, func(_ int, elem T) string {
		return fmt.Sprintf("%v", elem)
	}, elems...)
}

func JoinFn[T any](sep string, fn func(int, T) string, elems ...T) string {
	return gstring.Join(Map(fn, elems...), sep)
}

func JoinErrFn[T any](sep string, fn func(int, T) (string, error), elems ...T) (string, error) {
	res, err := MapErr(fn, elems...)

	if err != nil {
		return "", err
	}

	return gstring.Join(res, sep), nil
}
