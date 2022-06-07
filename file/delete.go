package file

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/code206/goutils/pathfunc"
)

// 删除文件，必须是文件才删除
func DeleteFile(elem ...string) error {
	var filePath string
	switch len(elem) {
	case 0:
		return errors.New("param is empty")
	case 1:
		filePath = elem[0]
	default:
		filePath = filepath.Join(elem...)
	}

	if filePath == "" {
		return errors.New("path is empty")
	}

	if pathfunc.IsFile(filePath) {
		return os.Remove(filePath)
	} else {
		return errors.New("path is not file")
	}
}
