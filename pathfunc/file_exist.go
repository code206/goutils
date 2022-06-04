package pathfunc

import "os"

// 判断所给路径是否为文件
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}
func FileExist(path string) bool {
	return IsFile(path)
}

// 判断所给路径是否不是文件
func FileNotExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return true
	}
	return fileInfo.IsDir()
}
