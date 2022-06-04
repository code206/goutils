package pathfunc

import (
	"os"
)

// 判断所给路径文件/文件夹是否存在。如果有链接，找到链接的目标，是否有真实的实体
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断所给路径文件/文件夹是否不存在。如果有链接，找到链接的目标，是否有真实的实体
func PathNotExist(path string) bool {
	return !PathExist(path)
}

// 路径是否存在，不跳转，只检查传入的 path 是否存在
func PathLinkExist(path string) bool {
	_, err := os.Lstat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 路径是否不存在，不跳转，只检查传入的 path 是否存在
func PathLinkNotExist(path string) bool {
	return !PathLinkExist(path)
}
