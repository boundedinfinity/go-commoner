package fileutil

import (
	"errors"
	"os"

	"github.com/boundedinfinity/go-commoner/try"
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

func PathExistsTry(path string) try.Try[bool] {
	return try.Complete(PathExistsErr(path))
}
