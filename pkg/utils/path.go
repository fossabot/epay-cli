package utils

import (
	"os"
	"path/filepath"
)

func GetAbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(exe)
	return filepath.Join(dir, path)
}
