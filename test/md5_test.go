package test

import (
	"fmt"
	"testing"
	"yutil/yfunc"
)

func TestMd5(t *testing.T){
	var a = "2d23d"
	fmt.Println(yfunc.Md5(a))
}

