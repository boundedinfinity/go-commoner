package pather

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

func PathExists(path string) trier.Try[bool] {
	return trier.Complete(pather.PathExistsErr(path))
}

func GetPaths(path string) trier.Try[[]string] {
	return trier.Complete(pather.GetPaths(path))
}
