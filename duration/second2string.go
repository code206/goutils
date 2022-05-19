package duration

import (
	"strconv"
	"time"
)

// 转换为秒级时长
func ToSecond(d time.Duration) int64 {
	return int64(d / time.Second)
}

// 秒级时长转为友好的文本描述
func SecondToString(sec int64) string {
	if day := sec / 86400; day > 0 {
		return strconv.FormatInt(day, 10) + "d"
	}
	if hour := sec / 3600; hour > 0 {
		return strconv.FormatInt(hour, 10) + "h"
	}
	if minute := sec / 60; minute > 0 {
		return strconv.FormatInt(minute, 10) + "m"
	}
	return strconv.FormatInt(sec, 10) + "s"
}
