package yutil

import (
	"errors"
	"strconv"
)

var versionErr = errors.New("does not support comparison")

// 比较两个版本号: v1 > v2 => 1; v1 == v2 => 0; v1 < v2 => -1
// 版本号示例: v1.0 v0.01 v1.0.0.0.1 v0.0.1 ...
func (yOther) CompareVersion(v1, v2 string) (int, error) {
	dot := false
	var first []byte
	for _, v := range []byte(v1) {
		if v == 46 && !dot {
			first = append(first, v)
			dot = true
		}
		if v >= 48 && v <= 58 {
			first = append(first, v)
		}
	}
	dot = false
	var second []byte
	for _, v := range []byte(v2) {
		if v == 46 && !dot {
			second = append(second, v)
			dot = true
		}
		if v >= 48 && v <= 58 {
			second = append(second, v)
		}
	}
	number1, err := strconv.ParseFloat(string(first), 64)
	if err != nil {
		return -2, versionErr
	}
	number2, err := strconv.ParseFloat(string(second), 64)
	if err != nil {
		return -3, versionErr
	}
	if number1 > number2 {
		return 1, nil
	} else if number1 < number2 {
		return -1, nil
	}
	return 0, nil
}
