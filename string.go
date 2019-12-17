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
