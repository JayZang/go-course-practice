package main

import (
	"fmt"
	"html/template"
	"math"
	"os"
)

// 要引入至template的函數要先存至template.FuncMap
var fm = template.FuncMap{
	"double": double,
	"square": square,
	"sqRoot": sqRoot,
}

func double(i int) int {
	return i * 2
}

func square(i int) float64 {
	return math.Pow(float64(i), 2)
}

func sqRoot(i float64) float64 {
	return math.Sqrt(i)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {
	data := 3
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}
}
