package pather

import (
	"io/fs"
	"os"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrNotFile  = errorer.Errorf("not a file")
	ErrNotFilev = ErrNotFile.ValueFn()
)

type FileConfig struct {
	Perm  os.FileMode
	Owner string
	Group string
}

var Files = files{
	config: FileConfig{
		Perm: 0755,
	},
}

type files struct {
	config FileConfig
}

func (t files) AbsErr(path string) (string, error) {
	return Paths.AbsErr(path)
}

func (t files) Abs(path string) string {
	return Paths.Abs(path)
}

func (t files) Exists(path string) bool {
	return Paths.Exists(path) && t.Is(path)
}

func (t files) ExistsErr(path string) (bool, error) {
	return Paths.ExistsErr(path)
}

func (t files) Is(path string) bool {
	ok, err := t.IsErr(path)

	if err != nil {
		panic(err)
	}

	return ok
}

func (t files) IsErr(path string) (bool, error) {
	info, err := os.Stat(string(path))

	if err != nil {
		return false, err
	}

	return info.Mode().IsRegular(), nil
}

func (t files) Remove(path string) bool {
	return Paths.Remove(path)
}

func (t files) RemoveErr(path string) (bool, error) {
	return Paths.RemoveErr(path)
}

func (t files) Dir(path string) string {
	return Paths.Dir(path)
}

func (t files) Ensure(path string) bool {
	ok, err := t.EnsureErr(path)

	if err != nil {
		panic(err)
	}

	return ok
}

func (t files) EnsureErr(path string) (bool, error) {
	exists, err := Paths.ExistsErr(path)

	if err != nil {
		return false, err
	}

	if exists {
		dir, err := t.IsErr(path)

		if err != nil {
			return false, err
		}

		if !dir {
			return false, ErrNotDirv(path)
		}
	} else {
		if err := os.MkdirAll(path, t.config.Perm); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (t files) Appendln(path string, bs []byte, mode fs.FileMode) error {
	return t.Append(path, append(bs, []byte("\n")...), mode)
}

func (t files) Append(path string, bs []byte, mode fs.FileMode) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)

	if err != nil {
		return err
	}

	defer file.Close()

	line := append(bs, []byte("\n")...)

	if _, err := file.Write(line); err != nil {
		panic(err)
	}

	return nil
}

func (t files) AppendlnString(path string, text string, mode fs.FileMode) error {
	return t.Appendln(path, []byte(text), mode)
}

func (t files) AppendString(path string, text string, mode fs.FileMode) error {
	return t.Append(path, []byte(text), mode)
}

func (t files) Walk(root string, filterFn func(string, fs.FileInfo) bool, processFn func(string, fs.FileInfo) error) error {
	filterFn2 := func(path string, info fs.FileInfo) bool {
		return !info.IsDir() && filterFn(path, info)
	}

	return Walks.Paths(root, filterFn2, processFn)
}

func (t files) List(root string) ([]string, error) {
	out := make([]string, 0)

	filterFn := func(path string, info os.FileInfo) bool {
		return info.Mode().IsRegular()
	}

	processFn := func(path string, info os.FileInfo) error {
		out = append(out, path)
		return nil
	}

	err := Walks.Paths(root, filterFn, processFn)

	return out, err
}
