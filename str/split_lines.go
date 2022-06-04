package str

import (
	"strings"
)

// 字符串转换为字符串切片，每行为一个元素
func SplitLines(str string) []string {
	s := strings.FieldsFunc(strings.TrimSpace(str), func(c rune) bool { return c == '\n' || c == '\r' })
	return s
}
