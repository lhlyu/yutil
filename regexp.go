package yutil

import "regexp"

// 正则切割字符串
func (yRegexp) SplitAll(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	return reg.Split(s, -1)
}

// 正则替换字符串
func (yRegexp) ReplaceAll(s, pattern string, repl string) string {
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(s, repl)
}

// 正则获取符合的所有字符串
func (yRegexp) FindAll(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	return reg.FindAllString(s, -1)
}
