package pather

import (
	"io/fs"
	"path/filepath"
)

var Walks = walks{}

type walks struct{}

func (t walks) Paths(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
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
