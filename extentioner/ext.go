package extentioner

import (
	"os"
	"path/filepath"
	"strings"
)

func Ext(path string) string {
	return filepath.Ext(path)
}

func ExtOnly(path string) string {
	return strings.Replace(filepath.Ext(path), ".", "", 1)
}

func Normalize(ext string) string {
	if strings.HasPrefix(ext, ".") {
		return ext
	} else {
		return "." + ext
	}
}

func Strip(path string) string {
	b := index(path)
	return path[:b]
}

func Swap(path, old, new string) string {
	old = Normalize(old)

	if !strings.HasSuffix(path, old) {
		return path
	}

	new = Normalize(new)
	ext := Ext(path)

	if ext == old {
		path = Strip(path)
		path = path + new
	}

	return path
}

func index(path string) int {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return i
		}
	}

	return -1
}
