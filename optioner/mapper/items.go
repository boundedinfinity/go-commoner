package mapper

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	"github.com/boundedinfinity/go-commoner/optioner"
)

func Items[K comparable, V any](m map[K]V) optioner.Option[[]mapper.Item[K, V]] {
	items := mapper.Items(m)

	if len(items) <= 0 {
		return optioner.None[[]mapper.Item[K, V]]()
	}

	return optioner.Some(items)
}
