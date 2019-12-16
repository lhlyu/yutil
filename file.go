package yutil

import "os"

// 判断文件是否存在
func FileIsExists(filePath string) bool{
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
