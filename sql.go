package yutil

import (
	"bytes"
	"fmt"
	"strings"
)

type SqlBuffer struct {
	buf bytes.Buffer
	ps  []interface{}
}

// sql 添加语句和参数
func (s *SqlBuffer) Add(text string, params ...interface{}) *SqlBuffer {
	s.buf.WriteString(text)
	s.ps = append(s.ps, params...)
	return s
}

func (s *SqlBuffer) AddWhr(text string, v interface{}) *SqlBuffer {
	s.buf.WriteString(text)
	s.ps = append(s.ps, v)
	return s
}

func (s *SqlBuffer) AddWhrNeqZero(text string, v int) *SqlBuffer {
	if v != 0 {
		s.buf.WriteString(text)
		s.ps = append(s.ps, v)
	}
	return s
}

func (s *SqlBuffer) AddWhrGtZero(text string, v int) *SqlBuffer {
	if v > 0 {
		s.buf.WriteString(text)
		s.ps = append(s.ps, v)
	}
	return s
}

func (s *SqlBuffer) AddWhrNeqEmpty(text string, v string) *SqlBuffer {
	if strings.TrimSpace(v) != "" {
		s.buf.WriteString(text)
		s.ps = append(s.ps, v)
	}
	return s
}

func (s *SqlBuffer) AddWhrIn(format string, params ...interface{}) *SqlBuffer {
	length := len(params)
	if length > 0 {
		text := fmt.Sprintf(format, CreateQuestionMarks(length))
		s.buf.WriteString(text)
		s.ps = append(s.ps, params...)
	}
	return s
}

func (s *SqlBuffer) AddWhrBatch(format string, v ...[]interface{}) *SqlBuffer {
	length := len(v)
	if length > 0 {
		text, params := CreateQuestionMarksForBatch(v...)
		text = fmt.Sprintf(format, text)
		s.buf.WriteString(text)
		s.ps = append(s.ps, params...)
	}
	return s
}

// 去除两侧的逗号
func (s *SqlBuffer) TrimComma() *SqlBuffer {
	if s.buf.Len() > 0 {
		str := s.buf.String()
		str = strings.TrimSpace(str)
		str = strings.Trim(str, ",")
		s.buf.Reset()
		s.buf.WriteString(str)
	}
	return s
}

func (s *SqlBuffer) Reset() *SqlBuffer {
	s.buf.Reset()
	s.ps = make([]interface{}, 0)
	return s
}

func (s *SqlBuffer) String() string {
	return strings.TrimSpace(s.buf.String())
}

func (s *SqlBuffer) Params() []interface{} {
	return s.ps
}

func (s *SqlBuffer) Build() (string, []interface{}) {
	return s.String(), s.Params()
}

// 根据长度创建占位符
func CreateQuestionMarks(length int) string {
	if length == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	buf.WriteString("?")
	buf.WriteString(strings.Repeat(",?", length-1))
	return buf.String()
}

// 批量添加时创建占位符
func CreateQuestionMarksForBatch(v ...[]interface{}) (string, []interface{}) {
	if len(v) == 0 {
		return "", nil
	}
	qNum := len(v[0])
	maxLength := len(v)
	qm := CreateQuestionMarks(qNum)
	var params []interface{}
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("(%s)", qm))
	params = append(params, v[0]...)
	for i := 1; i < maxLength; i++ {
		buf.WriteString(",(")
		buf.WriteString(qm)
		buf.WriteString(")")
		params = append(params, v[i]...)
	}
	return buf.String(), params
}
