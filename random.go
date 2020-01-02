package yutil

import (
	"math"
	"math/rand"
	"time"
)

// 生成区间的随机数(如果min参数大于max参数，这两个参数的值会对调)
func RandInterval(min, max int) int {
	if max < min {
		min, max = max, min
	} else if max == min {
		return max
	}
	time.Sleep(1)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(max-min) + min
	return x
}

// 随机一个数 1 << 63 -1
func RandNumber() int {
	time.Sleep(1)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(math.MaxInt64)
	return x
}

// 随机指定长度的字符串，(范围包含: 大小写数字下划线)
func RandString(length int) string {
	if length <= 0 {
		return ""
	}
	bs := []byte(ALPHA)
	bsLength := len(bs)
	rand.Seed(time.Now().UnixNano())
	var bts []byte
	for i := 0; i < length; i++ {
		time.Sleep(1)
		bts = append(bts, bs[rand.Intn(bsLength)])
	}
	return string(bts)
}

// 自定义随机字符串
func RandCustomString(length int, s string) string {
	if length <= 0 {
		return ""
	}
	bs := []byte(s)
	bsLength := len(bs)
	rand.Seed(time.Now().UnixNano())
	var bts []byte
	for i := 0; i < length; i++ {
		time.Sleep(1)
		bts = append(bts, bs[rand.Intn(bsLength)])
	}
	return string(bts)
}
