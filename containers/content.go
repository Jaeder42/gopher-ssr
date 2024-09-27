package containers

import (
	"bytes"
	"html/template"

	"jaeder42.gopherssr.com/generics"
)

type Content struct {
	generics.Component
	Child         generics.Component
	RenderedChild template.HTML
}

func NewContent(child generics.Component) Content {
	content := &Content{Child: child}
	return *content
}

func (c Content) Render() template.HTML {
	c.RenderedChild = c.Child.Render()
	tmpl, err := template.New("Content").Parse("<div>{{.RenderedChild}}</div>")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, c)
	if err != nil {
		panic(err)
	}
	return template.HTML(tpl.String())
}
