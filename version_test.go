package yutil

import (
	"fmt"
	"testing"
)

type Check struct {
	A string
	B string
	E int
}

func TestVersion(t *testing.T) {
	var checks = []Check{
		{"v1.0", "v1.0", 0},
		{"v1.0", "v1.0.0", 0},
		{"v1.0", "v1.00", 0},
		{"v1.0", "v1.0.0.0", 0},
		{"v1.0", "v1", 0},
		{"v1.0", "v1.1", -1},
		{"v1.0", "v1.0.1", -1},
		{"v1.0", "v1.01", -1},
		{"v1.0", "v1.0.0.0.1", -1},
		{"v0.9", "v0.88", 1},
		{"v0.9", "v0.9.1", -1},
		{"v1.0", "v0.999", 1},
		{"v1.0", "v0.8", 1},
		{"v1.0", "v0.7.7", 1},
		{"v1.0", "v0.0.1.0", 1},
		{"v1.0", "v0.5.0.1", 1},
		{"v32.0.1", "v0.5.0.1", 1},
		{"v1.0.1", "v1.1", -1},
		{"edfef", "dfe", -2},
	}
	for k, v := range checks {
		c, e := Other.CompareVersion(v.A, v.B)
		if e != nil {
			t.Log(e.Error())
			continue
		}
		fmt.Printf("%-2d > %-10s : %-10s = %-2d | %-2d | %v\n", k, v.A, v.B, c, v.E, c == v.E)
	}

}
