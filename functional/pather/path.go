package pather

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

var Paths = files{}

type paths struct{}

func (t paths) Exists(path string) trier.Try[bool] {
	return trier.Complete(pather.Paths.ExistsErr(path))
}

func (t paths) List(path string) trier.Try[[]string] {
	return trier.Complete(pather.Paths.List(path))
}

// func PathExists(path string) trier.Try[bool] {
// 	return trier.Complete(pather.PathExistsErr(path))
// }

// func GetPaths(path string) trier.Try[[]string] {
// 	return trier.Complete(pather.GetPaths(path))
// }
