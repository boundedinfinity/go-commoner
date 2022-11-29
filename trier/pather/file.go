package pather

import (
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/trier"
)

func IsFile(path string) trier.Try[bool] {
	return trier.Complete(pather.IsFile(path))
}

func GetFiles(path string) trier.Try[[]string] {
	return trier.Complete(pather.GetFiles(path))
}
