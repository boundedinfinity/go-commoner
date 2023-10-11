package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

func Items[K comparable, V any](m map[K]V) optioner.Option[[]mapper.Item[K, V]] {
	items := mapper.Items(m)

	if len(items) <= 0 {
		return optioner.None[[]mapper.Item[K, V]]()
	}

	return optioner.Some(items)
}
