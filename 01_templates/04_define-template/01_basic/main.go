package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template
var fm = template.FuncMap{}

func init() {
	// define之template必須和其他template檔案引入在同一個Template實例，否則會找不到define之template
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
}
