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
