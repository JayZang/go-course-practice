package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	handler := myHandler{}
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		panic(err)
	}
}

type myHandler struct{}

func (myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 設置 http Header
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("My-Header", "Hello World")

	tpl.ExecuteTemplate(w, "template.gohtml", nil)
}
