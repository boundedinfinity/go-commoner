package stringer

import "strings"

func Join[T any](sep string, elems ...T) string {
	strs := AsStrings(elems...)
	return strings.Join(strs, sep)
}

// func Join[T ~string](elems []T, sep string) string {
// 	s := slicer.Map(elems, func(elem T) string { return string(elem) })
// 	return strings.Join(s, sep)
// }

func JoinFunc[T any](elems []T, sep string, fn func(T) string) string {
	strs := SliceFunc(elems, fn)
	return strings.Join(strs, sep)
}

func SliceFunc[T any](elems []T, fn func(T) string) []string {
	output := make([]string, len(elems))

	for i, elem := range elems {
		output[i] = fn(elem)
	}

	return output
}
