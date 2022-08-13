package pathfunc

import (
	"path/filepath"
	"strings"
)

func TrimPathExt(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return path
	}
	return strings.TrimSuffix(path, ext)
}
