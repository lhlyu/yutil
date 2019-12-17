package yutil

import "strings"

// 过滤空字符
var FITER_EMPTY = func(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

// 过滤0
var FITER_ZERO = func(v int) bool {
	if v == 0 {
		return true
	}
	return false
}

// 过滤非正整数
var FITER_Z = func(v int) bool {
	if v <= 0 {
		return true
	}
	return false
}
