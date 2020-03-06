package yutil

import (
	"bytes"
	"html/template"
	"strings"
)

// 模板渲染
func (yString) TemplateParse(tmpl string, v interface{}) string {
	t := template.Must(template.New("template.html").Parse(tmpl))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, v)
	if err != nil {
		_arc.log("String", "TemplateParse", err)
		return ""
	}
	return buf.String()
}

// 字符串大驼峰
func (yString) BigCamelCase(s string) string {
	needCamel := true
	s = strings.Map(func(r rune) rune {
		if needCamel {
			if r == 32 || r == 95 {
				return -1
			}
			needCamel = false
			if r < 97 || r > 122 {
				return r
			}
			return r - 32
		}
		if r == 32 || r == 95 {
			needCamel = true
			return -1
		}
		return r
	}, s)
	return s
}

// 字符串小驼峰
func (yString) LittleCamelCase(s string) string {
	needCamel := false
	isFirst := true
	s = strings.Map(func(r rune) rune {
		if isFirst {
			isFirst = false
			if r < 65 || r > 90 {
				return r
			}
			return r + 32
		}
		if needCamel {
			if r == 32 || r == 95 {
				return -1
			}
			needCamel = false
			if r < 97 || r > 122 {
				return r
			}
			return r - 32
		}
		if r == 32 || r == 95 {
			needCamel = true
			return -1
		}
		return r
	}, s)
	return s
}
