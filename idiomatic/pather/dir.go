package pather

import (
	"io/fs"
	"os"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrNotDir  = errorer.Errorf("not a directory")
	ErrNotDirv = errorer.WithChild(ErrNotDir)
)

type DirConfig struct {
	Perm  os.FileMode
	Owner string
	Group string
}

func DirEnsure(path string) error {
	config := DirConfig{
		Perm: 0755,
	}

	return DirEnsureWithConfig(path, config)
}

func DirEnsureWithConfig(path string, config DirConfig) error {
	ok, err := PathExistsErr(path)

	if err != nil {
		return err
	}

	if !ok {
		if err := os.MkdirAll(path, config.Perm); err != nil {
			return err
		}
	}

	return nil
}

func IsDir[T ~string](path T) (bool, error) {
	info, err := os.Stat(string(path))

	if err != nil {
		return false, err
	}

	return info.Mode().IsDir(), nil
}

func WalkDirs(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
	filterFn2 := func(path string, info fs.FileInfo) bool {
		return info.IsDir() && filterFn(path, info)
	}

	return WalkPaths(root, filterFn2, processFn)
}

func GetDirs(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return true
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := WalkDirs(root, filterFn, processFn)

	return out, err
}
