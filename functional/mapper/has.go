package mapper

import "github.com/boundedinfinity/go-commoner/idiomatic/mapper"

func Has[K comparable, V any](m map[K]V, k K) bool {
	return mapper.Has(m, k)
}
