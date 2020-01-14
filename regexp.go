package yutil

import "regexp"

// v.1.0.6: 正则切割字符串
func RegexpSplitAll(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	return reg.Split(s, -1)
}

// v.1.0.6: 正则替换字符串
func RegexpReplaceAll(s, pattern string, repl string) string {
	reg := regexp.MustCompile(pattern)
	return string(reg.ReplaceAll([]byte(s), []byte(repl)))
}
