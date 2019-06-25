package test

import (
	"fmt"
	"testing"
	"yutil/yfunc"
)

func TestYio(t *testing.T) {
	path := "../README.md"
	//fmt.Println(yfunc.ReadFileAll(path))
	fmt.Println(len(yfunc.ReadFileLines(path)))
	fmt.Println(len(yfunc.ReadFileLinesTrim(path)))
}
