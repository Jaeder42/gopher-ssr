package generics

import "html/template"

type Component interface {
	Render() template.HTML
}
