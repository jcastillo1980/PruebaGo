package tools

import (
	"os"
)

func IsDirectory(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}

	return false
}
