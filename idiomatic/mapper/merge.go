package mapper

import "maps"

func Merge[K comparable, V any](dst, src map[K]V) {
	maps.Copy(dst, src)
}
