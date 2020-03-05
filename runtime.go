package yutil

import (
	"runtime"
)

const (
	OS_MAC     = "darwin"
	OS_WINDOWS = "windows"
	OS_LINUX   = "linux"
	OS_UNKNOWN = "unknown"
)

// 获取系统类型
func (yRuntime) GetOSName() string {
	switch runtime.GOOS {
	case OS_MAC:
		return OS_MAC
	case OS_WINDOWS:
		return OS_WINDOWS
	case OS_LINUX:
		return OS_LINUX
	default:
		return OS_UNKNOWN
	}
}

// 获取调用栈上的函数信息(函数名、文件名、调用行数)
func (yRuntime) CurrentFuncInfo(skip int) (funcName string, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	if ok {
		return f.Name(), file, line
	} else {
		return "", "", -1
	}
}
