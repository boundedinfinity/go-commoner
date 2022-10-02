package pather

import "os"

type DirConfig struct {
	Perm  os.FileMode
	Owner string
	Group string
}

func DirEnsure(path string) error {
	config := DirConfig{
		Perm: 0755,
	}

	return DirEnsureWithConfig(path, config)
}

func DirEnsureWithConfig(path string, config DirConfig) error {
	ok, err := PathExistsErr(path)

	if err != nil {
		return err
	}

	if !ok {
		if err := os.MkdirAll(path, config.Perm); err != nil {
			return err
		}
	}

	return nil
}

func IsDirErr(path string) (bool, error) {
	ok, err := IsFileErr(path)

	if err != nil {
		return false, err
	}

	return !ok, nil
}

func IsDir(path string) bool {
	ok, _ := IsDirErr(path)
	return ok
}
