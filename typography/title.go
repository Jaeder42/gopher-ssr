package typography

import (
	"bytes"
	"html/template"

	"jaeder42.gopherssr.com/generics"
)

type Title struct {
	generics.Component
	Text string
}

func NewTitle(text string) generics.Component {
	return &Title{
		Text: text,
	}
}
func (t Title) Render() template.HTML {
	tmpl, err := template.New("Text").Parse("<h1>{{.Text}}</h1>")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, t)
	if err != nil {
		panic(err)
	}
	return template.HTML(tpl.String())
}
