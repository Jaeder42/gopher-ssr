package typography

import (
	"bytes"
	"html/template"

	"jaeder42.gopherssr.com/generics"
)

type Text struct {
	generics.Component
	Text string
}

func NewText(text string) generics.Component {
	return Text{Text: text}
}

func (t Text) Render() template.HTML {
	tmpl, err := template.New("Text").Parse("<p>{{.Text}}</p>")
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
