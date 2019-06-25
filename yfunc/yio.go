package yfunc

import (
	"bufio"
	"github.com/lhlyu/logger"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// @一次性读取文件所有内容
func ReadFileAll(filePath string) string{
	fi, err := os.Open(filePath)
	if err != nil {
		logger.Error("filePath is err : ",err.Error())
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		logger.Error("ioutil.ReadAll is err : ",err.Error())
		return ""
	}
	return string(fd)
}

// @逐行读取文件内容
func ReadFileLines(filePath string) []string{
	var lines []string
	fi, err := os.Open(filePath)
	if err != nil {
		logger.Error("filePath is err : ",err.Error())
		return lines
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		line, _,err := br.ReadLine()
		if err == io.EOF{
			break
		}
		lines = append(lines,string(line))
	}
	return lines
}

// @逐行读取文件内容，去除空行和首尾空格
func ReadFileLinesTrim(filePath string) []string{
	var lines []string
	fi, err := os.Open(filePath)
	if err != nil {
		logger.Error("filePath is err : ",err.Error())
		return lines
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		linebyt, _,err := br.ReadLine()
		if err == io.EOF{
			break
		}
		line := string(linebyt)
		line = strings.TrimSpace(line)
		if line == ""{
			continue
		}
		lines = append(lines,line)
	}
	return lines
}
