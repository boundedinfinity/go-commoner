package pather

import (
	"io/fs"
	"os"
)

func IsFileErr(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.Mode().IsRegular(), nil
}

func IsFile(path string) bool {
	ok, _ := IsFileErr(path)
	return ok
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
