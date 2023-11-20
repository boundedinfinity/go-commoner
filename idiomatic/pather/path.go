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

func (t paths) Remove(path string) {
	t.RemoveErr(path)
}

func (t paths) RemoveErr(path string) error {
	ok, err := t.ExistsErr(path)

	if err != nil {
		return err
	}

	if ok {
		return os.Remove(path)
	}

	return nil
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

// func Base[T ~string](path T) string {
// 	return filepath.Base(string(path))
// }

// func Dir[T ~string](path T) string {
// 	return filepath.Dir(string(path))
// }

// func PathEnsureRemove(path string) {
// 	PathExistsErr(path)
// }

// func PathEnsureRemoveErr(path string) error {
// 	ok, err := PathExistsErr(path)

// 	if err != nil {
// 		return err
// 	}

// 	if ok {
// 		return os.Remove(path)
// 	}

// 	return nil
// }

// func PathExists(path string) bool {
// 	exist, _ := PathExistsErr(path)
// 	return exist
// }

// func PathExistsErr(path string) (bool, error) {
// 	_, err := os.Stat(path)

// 	if err != nil {
// 		if errors.Is(err, os.ErrNotExist) {
// 			return false, nil
// 		} else {
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }

// func GetPaths(root string) ([]string, error) {
// 	out := make([]string, 0)

// 	filterFn := func(path string, info os.FileInfo) bool {
// 		return true
// 	}

// 	processFn := func(path string, info os.FileInfo) error {
// 		out = append(out, path)
// 		return nil
// 	}

// 	err := WalkPaths(root, filterFn, processFn)

// 	return out, err
// }
