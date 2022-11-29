package mapper

type Mapper[K comparable, V any] map[K]V

func (t Mapper[K, V]) Has(k K) bool {
	return Has(t, k)
}

func (t Mapper[K, V]) Keys() []K {
	return Keys(t)
}

func (t Mapper[K, V]) KeysFiltered(fn func(K) bool) []K {
	return KeysFiltered(t, fn)
}

func (t Mapper[K, V]) Values() []V {
	return Values(t)
}

func (t Mapper[K, V]) ValuesFiltered(fn func(V) bool) []V {
	return ValuesFiltered(t, fn)
}

func (t Mapper[K, V]) Items() []Item[K, V] {
	return Items(t)
}

func (t Mapper[K, V]) Each(fn func(K, V)) {
	Each(t, fn)
}

func (t Mapper[K, V]) EachErr(fn func(K, V) error) error {
	return EachErr(t, fn)
}
