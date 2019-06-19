package test

import (
	"fmt"
	"testing"
	"yutil/yfunc"
)

func TestBase64(t *testing.T){
	s := "e323e哈#@*("
	encode := yfunc.AnyEncode(s)
	fmt.Println("encode = ",encode)
	var decode string
	yfunc.AnyDecode(encode,&decode)
	fmt.Println("decode = ",decode)
}
