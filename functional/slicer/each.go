package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func Each[T any](fn func(int, T), elems ...T) {
	slicer.Each(fn, elems...)
}

func EachErr[T any](fn func(int, T) error, elems ...T) error {
	return slicer.EachErr(fn, elems...)
}
