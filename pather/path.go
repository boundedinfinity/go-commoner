package pather

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func PathExists(path string) bool {
	exist, _ := PathExistsErr(path)
	return exist
}

func PathExistsErr(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func WalkPaths(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
	walkFn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !filterFn(path, info) {
			return nil
		}

		if err := processFn(path, info); err != nil {
			return err
		}

		return nil
	}

	return filepath.Walk(root, walkFn)
}

func GetPaths(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return true
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := WalkPaths(root, filterFn, processFn)

	return out, err
}
