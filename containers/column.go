package containers

import (
	"bytes"
	"html/template"

	"jaeder42.gopherssr.com/generics"
)

type Column struct {
	generics.Component
	Children         []generics.Component
	RenderedChildren template.HTML
}

func NewColumn(children []generics.Component) generics.Component {
	return &Column{Children: children}

}

func (c Column) Render() template.HTML {
	for _, child := range c.Children {
		c.RenderedChildren += child.Render()
	}
	tmpl, err := template.New("Text").Parse("<div style='display: flex; flex-direction: column;'>{{.RenderedChildren}}</div>")
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
