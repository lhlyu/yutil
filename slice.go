package yutil

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// v1.0.0: 将字符串切片转成int切片
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

// v1.0.0: 将int切片转成字符串切片
func SliceIntToStr(vv []int) []string {
	var slice []string
	for _, v := range vv {
		val := strconv.Itoa(v)
		slice = append(slice, val)
	}
	return slice
}

// v1.0.0: 过滤字符串切片,如果 f 方法返回true 则过滤当前元素
func SliceStrFilter(ss []string, f func(s string) bool) {
	var slice []string
	for _, s := range ss {
		if !f(s) {
			slice = append(slice, s)
		}
	}
}

// v1.0.0: 过滤int切片,如果 f 方法返回true 则过滤当前元素
func SliceIntFilter(vv []int, f func(v int) bool) {
	var slice []int
	for _, v := range vv {
		if !f(v) {
			slice = append(slice, v)
		}
	}
}

// v1.0.7: 切片洗牌
func SliceShuffle(v interface{}){
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr{
		val = val.Elem()
	}
	if val.Kind() != reflect.Slice{
		return
	}
	length := val.Len()
	rand.Seed(time.Now().UnixNano())
	for i := length - 1; i > 0;i --{
		j := rand.Intn(length)
		reflect.Swapper(v)(i,j)
	}
}