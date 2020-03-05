package yutil

import (
	"math/rand"
	"time"
)

// 随机指定长度的字符串，(范围包含: 大小写数字下划线)
func (yRandom) String(length int) string {
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
