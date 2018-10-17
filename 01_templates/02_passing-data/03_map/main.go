package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	people := map[string]int{
		"Jay":   18,
		"Ban":   30,
		"Alice": 40,
	}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		fmt.Println(err)
	}
}
