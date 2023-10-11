package pather

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

func IsDir(path string) trier.Try[bool] {
	return trier.Complete(pather.IsDir(path))
}

func GetDirs(path string) trier.Try[[]string] {
	return trier.Complete(pather.GetDirs(path))
}
