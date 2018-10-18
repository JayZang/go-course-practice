package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

// 要引入至template的函數要先存至template.FuncMap
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) (out string) {
	out = strings.TrimSpace(s)
	out = out[:3]
	return
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {
	data := "Hello"
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}
}
