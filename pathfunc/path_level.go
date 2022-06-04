package pathfunc

import (
	"errors"
	"os"
)

// abcdefghijk  1:2:3  a/bc/def 1位有16个，2位256个
func PathLevels(str string, levels []int) (string, error) {
	var (
		subPath string
		start   int
		l       int = len(str)
	)

	for _, i := range levels {
		if l < start+i {
			return "", errors.New("string is too shoot")
		}

		subPath = subPath + string(os.PathSeparator) + str[start:start+i]
		start = start + i
	}

	return subPath[1:], nil
}
