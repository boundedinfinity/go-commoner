package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func AnyOf[S ~string, E ~string](s S, elems ...E) bool {
	var temp []string

	for _, elem := range elems {
		temp = append(temp, string(elem))
	}

	return slicer.AnyOf(string(s), temp...)
}

func AnyOfIgnoreCase[S ~string, E ~string](s S, elems ...E) bool {
	return AnyOf(Lowercase(s), LowercaseAll(elems...)...)
}
