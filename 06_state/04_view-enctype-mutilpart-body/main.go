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
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		b := make([]byte, req.ContentLength)
		req.Body.Read(b)
		body := string(b)

		tpl.Execute(res, body)
	})

	http.ListenAndServe(":8080", nil)
}
