package test

import (
	"testing"
	"yutil/yfunc"
)

func TestVerision(t *testing.T){
		var fibTests = []struct {
			v1 string
			v2 string
			expected int // expected result
		}{
			{"v1.0.0","v1.0.1",-1},
			{"v1.0.0","v1.0.2",-1},
			{"v1.0.0","v1.1.1",-1},
			{"v1.0.0","v1.1.0",-1},
			{"v1.0.0","v2.0.0",-1},
			{"v1.0.0","v1.2.0.0",-1},
			{"v1.0.1","v1.0.1a",-1},
			{"v0.0.1","v1.0.1",-1},
			{"v0.0.0","v0.0.0",0},
			{"v0.0.0","v0",0},
			{"v1.0.0","v1.0",0},
			{"v1.0.0","v1.0.0.0",0},
			{"v2.0.0","v1.0.0",1},
			{"v2.1.0","v1.0.1",1},
			{"v1.2.3","v1.1.1",1},
		}

		for _, tt := range fibTests {
			real := yfunc.CompareVersion(tt.v1,tt.v2)
			if real != tt.expected {
				t.Errorf("v1=%s\tv2=%s\t\t\texpected=%d\treal=%d", tt.v1,tt.v2, tt.expected,real)
			}
		}
}
