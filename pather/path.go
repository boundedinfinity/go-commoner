package pather

import (
	"errors"
	"os"

	"github.com/boundedinfinity/go-commoner/try"
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

func PathExistsTry(path string) try.Try[bool] {
	return try.Complete(PathExistsErr(path))
}
