package yutil

import (
	"math/rand"
	"reflect"
	"time"
)

// 切片洗牌
func (ySlice) Shuffle(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Slice {
		return
	}
	length := val.Len()
	rand.Seed(time.Now().UnixNano())
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(length)
		reflect.Swapper(v)(i, j)
	}
}

// 切片去重,如果传入指针则直接操作指针,如果传入非指针则返回一个interface，可自行断言
func (ySlice) Distinct(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	isPtr := false
	if val.Kind() == reflect.Ptr {
		isPtr = true
		val = val.Elem()
	}
	if val.Kind() != reflect.Slice {
		return v
	}
	existMap := make(map[interface{}]bool)
	outArr := reflect.MakeSlice(val.Type(), 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		iVal := val.Index(i)

		if _, ok := existMap[iVal.Interface()]; !ok {
			outArr = reflect.Append(outArr, val.Index(i))
			existMap[iVal.Interface()] = true
		}
	}
	if isPtr {
		val.Set(outArr)
	}
	return outArr.Interface()
}
