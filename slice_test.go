package yutil

import (
	"fmt"
	"testing"
)

func TestDistinct(t *testing.T) {
	a := []int{1, 2, 3, 1, 2, 4, 5, 4, 6, 1}
	Slice.Distinct(&a)
	fmt.Println(a)
	b := []string{"1", "2", "3", "4", "5", "1", "2", "1"}
	b = Slice.Distinct(b).([]string)
	fmt.Println(b)
}

func TestYSlice_Shuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	Slice.Shuffle(a)
	fmt.Println(a)
}
