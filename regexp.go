package yutil

import "regexp"

// v1.0.6: 正则切割字符串
func RegexpSplitAll(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	return reg.Split(s, -1)
}

// v1.0.6: 正则替换字符串
// v1.0.7: update
func RegexpReplaceAll(s, pattern string, repl string) string {
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(s,repl)
}

// v1.0.7: 正则获取符合的所有字符串
func RegexpFindAll(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	return reg.FindAllString(s,-1)
}