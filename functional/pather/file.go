package pather

import (
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

type FileConfig = pather.FileConfig

var Files = files{
	config: FileConfig{
		Perm: 0755,
	},
}

type files struct {
	config FileConfig
}

func (t files) Is(path string) trier.Try[bool] {
	return trier.CompleteErr(pather.Files.IsErr(path))
}

func (t files) List(path string) trier.Try[[]string] {
	return trier.CompleteErr(pather.Files.List(path))
}

// func IsFile(path string) trier.Try[bool] {
// 	return trier.Complete(pather.IsFile(path))
// }

// func GetFiles(path string) trier.Try[[]string] {
// 	return trier.Complete(pather.GetFiles(path))
// }
