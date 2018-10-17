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
	// 將 data 傳入 template
	err := tpl.Execute(os.Stdout, "Hello")
	if err != nil {
		fmt.Println(err)
	}
}
