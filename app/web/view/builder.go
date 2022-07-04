package view

import (
	"fmt"
	"github.com/gin-gonic/gin/render"
	"html/template"
)

var tpl = map[string]string{
	"/model": model,
}

func Build(tplName string, data any) (html render.HTML, err error) {
	if str, ok := tpl[tplName]; ok {
		v := template.New("/model")
		if v, err = v.Parse(str); err != nil {
			return html, err
		}
		html.Template = v
		html.Data = data
		return
	}
	return html, fmt.Errorf("cannot found %s tpl", tplName)
}
