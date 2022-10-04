package slicer

import (
	"fmt"
	gstring "strings"
)

func Join[T any](items []T, sep string) string {
	return JoinFn(items, sep, func(el T) string {
		return fmt.Sprintf("%v", el)
	})
}

func JoinFn[T any](items []T, sep string, fn func(T) string) string {
	return gstring.Join(Map(items, fn), sep)
}

func JoinErrFn[T any](items []T, sep string, fn func(T) (string, error)) (string, error) {
	res, err := MapErr(items, fn)

	if err != nil {
		return "", err
	}

	return gstring.Join(res, sep), nil
}
