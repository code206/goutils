package http

import "strings"

// 通过 uri，解析 query 和 path 获取参数，注意：建议只对处理好的 uri 做解析
// arg1/10/arg2/abc 或者 /path?arg1=10&arg2=abc
// 解析返回为 map{arg1:"10", arg2:"abc"}
func ParseUriArgs(uri string) map[string]string {
	argsMap := map[string]string{}
	path, query, found := strings.Cut(uri, "?")

	if found && query != "" {
		args := strings.Split(query, "&")
		for i := range args {
			key, value, found := strings.Cut(args[i], "=")
			if found && value != "" {
				argsMap[key] = value
			}
		}
		return argsMap
	}

	p := strings.TrimPrefix(path, "/")
	args := strings.Split(p, "/")
	argsLen := len(args)
	for i := 0; i < argsLen; i = i + 2 {
		if i+1 > argsLen {
			break
		}
		argsMap[args[i]] = args[i+1]
	}
	return argsMap
}
