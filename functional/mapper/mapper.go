package mapper

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

type Mapper[K comparable, V any] mapper.Mapper[K, V]

func New[K comparable, V any]() Mapper[K, V] {
	return Mapper[K, V]{}
}
