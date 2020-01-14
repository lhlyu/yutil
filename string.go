package yutil

import (
	"bytes"
	"html/template"
)

// 模板渲染
func TemplateParse(tmpl string, v interface{}) string {
	t := template.Must(template.New("template.html").Parse(tmpl))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, v)
	if err != nil {
		_conf.log("string", "TemplateParse", err)
		return ""
	}
	return buf.String()
}

// v1.0.2：比较两个字符串，如果不相等，返回第二个
// 如果相等并且第二个参数不为空，则返回第一个
func CompareStr(first, second string) string {
	if second == "" {
		return first
	}
	if first != second {
		return second
	}
	return first
}

// v1.0.2: 比较两个整数，如果不相等，返回第二个
func CompareInt(first, second int) int {
	if first != second {
		return second
	}
	return first
}

// v1.0.2: 比较两个正整数，如果不相等，返回第二个
func CompareZInt(first, second int) int {
	if second <= 0 {
		return first
	}
	if first != second {
		return second
	}
	return first
}
