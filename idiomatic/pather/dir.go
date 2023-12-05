package pather

import (
	"io/fs"
	"os"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrNotDir  = errorer.Errorf("not a directory")
	ErrNotDirv = errorer.Subf(ErrNotDir)
)

type DirConfig struct {
	Perm  os.FileMode
	Owner string
	Group string
}

var Dirs = dirs{
	config: DirConfig{
		Perm: 0755,
	},
}

type dirs struct {
	config DirConfig
}

func (t dirs) Config(config DirConfig) dirs {
	return dirs{config}
}

func (t dirs) Ensure(path string) bool {
	ok, _ := t.EnsureErr(path)
	return ok
}

func (t dirs) EnsureErr(path string) (bool, error) {
	ok, err := Paths.ExistsErr(path)

	if err != nil {
		return false, err
	}

	if !ok {
		if err := os.MkdirAll(path, t.config.Perm); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (t dirs) Is(path string) bool {
	ok, _ := t.IsErr(path)
	return ok
}

func (t dirs) IsErr(path string) (bool, error) {
	info, err := os.Stat(string(path))

	if err != nil {
		return false, err
	}

	return info.Mode().IsDir(), nil
}

func (t dirs) AbsErr(path string) (string, error) {
	return Paths.AbsErr(path)
}

func (t dirs) Abs(path string) string {
	return Paths.Abs(path)
}

func (t dirs) Exists(path string) bool {
	return Paths.Exists(path)
}

func (t dirs) ExistsErr(path string) (bool, error) {
	return Paths.ExistsErr(path)
}

func (t dirs) Remove(path string) bool {
	return Paths.Remove(path)
}

func (t dirs) RemoveErr(path string) (bool, error) {
	return Paths.RemoveErr(path)
}

func (t dirs) Add(dst string, src ...string) error {
	for _, s := range src {
		if err := os.Rename(s, dst); err != nil {
			return err
		}
	}

	return nil
}

func (t dirs) Walk(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
	filterFn2 := func(path string, info fs.FileInfo) bool {
		return info.IsDir() && filterFn(path, info)
	}

	return Walks.Paths(root, filterFn2, processFn)
}

func (t dirs) List(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return true
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := Dirs.Walk(root, filterFn, processFn)

	return out, err
}
