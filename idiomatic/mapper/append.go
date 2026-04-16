package mapper

import "maps"

func Append[K comparable, V any](srcs ...map[K]V) map[K]V {
	dst := make(map[K]V)

	for _, src := range srcs {
		maps.Copy(dst, src)
	}

	return dst
}
