package pather

import (
	"io/fs"
	"path/filepath"
)

func WalkFilePaths(root string, filterFn func(string) bool, processFn func(string) error) error {
	wrapper := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !filterFn(path) {
			return nil
		}

		if err := processFn(path); err != nil {
			return err
		}

		return nil
	}

	return filepath.Walk(root, wrapper)
}

func WalkDirPaths(root string, filter func(string) bool, processFn func(string) error) error {
	wrapper := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if !filter(path) {
			return nil
		}

		if err := processFn(path); err != nil {
			return err
		}

		return nil
	}

	return filepath.Walk(root, wrapper)
}
