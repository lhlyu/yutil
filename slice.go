package yutil

import (
	"strconv"
)

// 将字符串切片转成int切片
func SliceStrToInt(ss []string) []int {
	var slice []int
	for _, v := range ss {
		val, err := strconv.Atoi(v)
		_conf.log("slice", "SliceStrToInt", err)
		if err != nil {
			continue
		}
		slice = append(slice, val)
	}
	return slice
}

// 将int切片转成字符串切片
func SliceIntToStr(vv []int) []string {
	var slice []string
	for _, v := range vv {
		val := strconv.Itoa(v)
		slice = append(slice, val)
	}
	return slice
}

// 过滤字符串切片,如果 f 方法返回true 则过滤当前元素
func SliceStrFilter(ss []string, f func(s string) bool) {
	var slice []string
	for _, s := range ss {
		if !f(s) {
			slice = append(slice, s)
		}
	}
}

// 过滤int切片,如果 f 方法返回true 则过滤当前元素
func SliceIntFilter(vv []int, f func(v int) bool) {
	var slice []int
	for _, v := range vv {
		if !f(v) {
			slice = append(slice, v)
		}
	}
}
