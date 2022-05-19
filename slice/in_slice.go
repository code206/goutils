package slice

import (
	"strings"
)

// 必须完全相同
func InSlice(need string, needSlice []string) bool {
	for _, v := range needSlice {
		if need == v {
			return true
		}
	}
	return false
}

// 不区分大小写
func InSliceEqualFold(need string, needSlice []string) bool {
	for _, v := range needSlice {
		if strings.EqualFold(need, v) {
			return true
		}
	}
	return false
}

// 前缀相同
func InSliceHasPrefix(s string, prefixs []string) bool {
	for _, prefix := range prefixs {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// 后缀相同
func InSliceHasSuffix(s string, suffixs []string) bool {
	for _, suffix := range suffixs {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}
