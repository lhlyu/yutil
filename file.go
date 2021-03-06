package yutil

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 判断文件是否存在
func (yFile) IsExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 一次性读取文件所有内容
func (yFile) ReadAll(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		_arc.log("File", "ReadAll", err)
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		_arc.log("File", "ReadAll", err)
		return ""
	}
	return string(fd)
}

// 按行读取
func (yFile) ReadLines(filePath string) []string {
	var lines []string
	fi, err := os.Open(filePath)
	if err != nil {
		_arc.log("File", "ReadLines", err)
		return lines
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	return lines
}

// 逐行读取文件内容，去除空行和首尾空格
func (yFile) ReadLinesTrim(filePath string) []string {
	var lines []string
	fi, err := os.Open(filePath)
	if err != nil {
		_arc.log("File", "ReadLinesTrim", err)
		return lines
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		linebyt, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(linebyt)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

// 自定义处理每一行
func (yFile) ReadLine(filePath string, f func(line string)) {
	fi, err := os.Open(filePath)
	if err != nil {
		_arc.log("File", "ReadLine", err)
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		f(string(line))
	}
}
