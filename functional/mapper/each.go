package mapper

import "github.com/boundedinfinity/go-commoner/idiomatic/mapper"

func Each[K comparable, V any](m map[K]V, fn func(K, V)) {
	mapper.Each(m, fn)
}

func EachErr[K comparable, V any](m map[K]V, fn func(K, V) error) error {
	return mapper.EachErr(m, fn)
}
