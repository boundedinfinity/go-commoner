package fileutil

import (
	"path/filepath"
	"strings"
)

func NormalizeExt(ext string) string {
	if strings.HasPrefix(ext, ".") {
		return ext
	} else {
		return "." + ext
	}
}

func SwapExt(path, oldExt, newExt string) string {
	oldExt2 := NormalizeExt(oldExt)

	if !strings.HasSuffix(path, oldExt) {
		return path
	}

	newExt2 := NormalizeExt(newExt)
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	filename = strings.ReplaceAll(filename, oldExt2, newExt2)
	newPath := filepath.Join(dir, filename)
	return newPath
}
