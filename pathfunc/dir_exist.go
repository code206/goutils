package pathfunc

import (
	"errors"
	"os"
	"path/filepath"
)

// 判断所给路径是否为文件夹。如果有链接，找到链接的目标，是否有真实的实体
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// 判断所给路径是否存在子目录
func HasSubDir(path string) (bool, error) {
	if !IsDir(path) {
		return false, errors.New("path is not dir")
	}
	//获取当前目录下的文件或目录名(包含路径)
	filespathNames, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return false, err
	}

	for i := range filespathNames {
		if IsDir(filespathNames[i]) {
			return true, nil
		}
	}

	return false, nil
}
