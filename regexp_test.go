package yutil

import (
	"fmt"
	"testing"
)

func TestRegexpSplitAll(t *testing.T) {
	s := "2s2s3d4gfd54g"
	arr := RegexpSplitAll(s, "\\d+")
	for k, v := range arr {
		t.Log(k, v)
	}
}

func TestRegexpReplaceAll(t *testing.T) {
	s := "s2s2s3d4gfd54g"
	s = RegexpReplaceAll(s, "\\d+|g", "_")
	fmt.Println(s)
}
