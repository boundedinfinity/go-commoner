package pather

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

type DirConfig = pather.DirConfig

var Dirs = dirs{}

type dirs struct{}

func (t dirs) Is(path string) trier.Try[bool] {
	return trier.CompleteErr(pather.Dirs.IsErr(path))
}

func (t dirs) List(path string) trier.Try[[]string] {
	return trier.CompleteErr(pather.Dirs.List(path))
}

// func IsDir(path string) trier.Try[bool] {
// 	return trier.Complete(pather.IsDir(path))
// }

// func GetDirs(path string) trier.Try[[]string] {
// 	return trier.Complete(pather.GetDirs(path))
// }
