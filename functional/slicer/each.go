package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func Each[T any](fn func(T), items ...T) {
	slicer.Each(fn, items...)
}

func EachErr[T any](fn func(T) error, items ...T) error {
	return slicer.EachErr(fn, items...)
}
