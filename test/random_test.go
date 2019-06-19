package test

import (
	"fmt"
	"testing"
	"yutil/yfunc"
)

func TestRandom(t *testing.T){
	fmt.Println(yfunc.GetRandomString(8))
}