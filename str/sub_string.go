package str

// rune字符截取
func SubString(s string, start, end int) string {
	sr := []rune(s)
	srLen := len(sr)

	// 如果在范围内，直接返回
	if start == 0 && srLen <= end {
		return s
	}

	if srLen == 0 || srLen <= start || end <= start {
		return ""
	}

	if srLen <= end {
		return string(sr[start:])
	}

	return string(sr[start:end])
}
