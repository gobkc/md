package global

import (
	"bytes"
	"html/template"
)

//解析和替换HTML中的内容
func ParseHtml(content string, replace interface{}) (result string, err error) {
	var c *template.Template
	if c, err = template.New("member").Parse(content); err != nil {
		return result, err
	}
	var buf bytes.Buffer
	c.Execute(&buf, replace)
	result = buf.String()
	return result, err
}
