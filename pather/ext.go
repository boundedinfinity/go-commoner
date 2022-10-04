package pather

import (
	"path/filepath"
	"strings"
)

func Ext(path string) string {
	return filepath.Ext(path)
}

func NormalizeExt(ext string) string {
	if strings.HasPrefix(ext, ".") {
		return ext
	} else {
		return "." + ext
	}
}

func SwapExt(path, oldExt, newExt string) string {
	normOldExt := NormalizeExt(oldExt)

	if !strings.HasSuffix(path, oldExt) {
		return path
	}

	normNewExt := NormalizeExt(newExt)
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	filename = strings.ReplaceAll(filename, normOldExt, normNewExt)
	newPath := filepath.Join(dir, filename)
	return newPath
}
