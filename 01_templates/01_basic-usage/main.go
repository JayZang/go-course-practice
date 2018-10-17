package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

// 執行main函數前或當作package被import時，init函數皆會先被執行
func init() {
	// 讀入tempaltes，可視為templates container
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	// 輸出指定的template
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gotext", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "bye.gotext", nil)
	if err != nil {
		fmt.Println(err)
	}

	// 未指定時預設為第一個讀到的template
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}
