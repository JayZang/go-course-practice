package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

// 成員變數使用大寫，外部package才可以讀取操作
type People struct {
	Name string
	Old  int
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	people := People{
		"Jay",
		18,
	}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		fmt.Println(err)
	}
}
