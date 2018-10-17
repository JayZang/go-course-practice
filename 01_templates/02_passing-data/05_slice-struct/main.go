package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

type People struct {
	Name string
	Old  int
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	jay := People{
		"Jay",
		18,
	}

	bob := People{
		"Bob",
		22,
	}

	people := []People{
		jay,
		bob,
	}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		fmt.Println(err)
	}
}
