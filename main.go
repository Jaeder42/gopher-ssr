package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"jaeder42.gopherssr.com/containers"
	"jaeder42.gopherssr.com/generics"
	"jaeder42.gopherssr.com/typography"
)

type Body struct {
	Body template.HTML
}

func main() {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		panic(err)
	}
	content := containers.NewColumn(
		[]generics.Component{
			typography.NewTitle("Testa"),
			typography.NewText("Test!")})

	var body Body
	var tpl bytes.Buffer
	body.Body = content.Render()
	if err := t.Execute(&tpl, &body); err != nil {
		panic(err)
	}

	result := tpl.String()
	error := os.WriteFile("./tmp/out.html", []byte(result), 0644)
	if error != nil {
		panic(error)
	}

	fmt.Println(result)

}
