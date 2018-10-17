package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

type Home struct {
	Member []People
	Floor  int
}

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

	home := Home{
		Member: []People{
			jay,
			bob,
		},
		Floor: 3,
	}

	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		fmt.Println(err)
	}
}
