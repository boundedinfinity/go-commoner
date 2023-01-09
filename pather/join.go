package pather

import (
	"path/filepath"
)

func Join[T ~string](elems ...T) string {
	var s []string

	for _, elem := range elems {
		s = append(s, string(elem))
	}

	return filepath.Join(s...)
}
