package pather

import (
	"errors"
	"os"
	"path/filepath"
)

var Paths = paths{}

type paths struct{}

func (t paths) Base(path string) string {
	return filepath.Base(string(path))
}

func (t paths) Dir(path string) string {
	return filepath.Dir(string(path))
}

func (t paths) AbsErr(path string) (string, error) {
	return filepath.Abs(path)
}

func (t paths) Abs(path string) string {
	abs, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return abs
}

func (t paths) Exists(path string) bool {
	exist, _ := t.ExistsErr(path)
	return exist
}

func (t paths) ExistsErr(path string) (bool, error) {
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

func (t paths) Remove(path string) bool {
	ok, err := t.RemoveErr(path)

	if err != nil {
		panic(err)
	}

	return ok
}

func (t paths) RemoveErr(path string) (bool, error) {
	ok, err := t.ExistsErr(path)

	if err != nil {
		return false, err
	}

	if ok {
		err = os.Remove(path)
	}

	if err != nil {
		return false, os.Remove(path)
	}

	return ok, nil
}

func (t paths) List(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return true
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := Walks.Paths(root, filterFn, processFn)

	return out, err
}
