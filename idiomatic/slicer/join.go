package slicer

import (
	"fmt"
	gstring "strings"
)

func Join[T any](sep string, items ...T) string {
	return JoinFn(sep, func(el T) string {
		return fmt.Sprintf("%v", el)
	}, items...)
}

func JoinFn[T any](sep string, fn func(T) string, items ...T) string {
	return gstring.Join(Map(fn, items...), sep)
}

func JoinErrFn[T any](sep string, fn func(T) (string, error), items ...T) (string, error) {
	res, err := MapErr(fn, items...)

	if err != nil {
		return "", err
	}

	return gstring.Join(res, sep), nil
}
