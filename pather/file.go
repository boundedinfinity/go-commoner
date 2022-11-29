package pather

import (
	"io/fs"
	"os"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrNotFile  = errorer.Errorf("not a file")
	ErrNotFilev = errorer.Errorfn(ErrNotDir)
)

func IsFileErr[T ~string](path T) error {
	info, err := os.Stat(string(path))

	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return ErrNotFilev(path)
	}

	return nil
}

func IsFile[T ~string](path T) (bool, error) {
	info, err := os.Stat(string(path))

	if err != nil {
		return false, err
	}

	return info.Mode().IsRegular(), nil
}

func WalkFiles(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
	filterFn2 := func(path string, info fs.FileInfo) bool {
		return !info.IsDir() && filterFn(path, info)
	}

	return WalkPaths(root, filterFn2, processFn)
}

func GetFiles(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return true
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := WalkFiles(root, filterFn, processFn)

	return out, err
}
