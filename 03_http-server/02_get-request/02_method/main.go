package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	handler := myHandler{}

	// 建立 tcp server，應用層為 http
	http.ListenAndServe(":8080", handler)
}

// 實作 Handler Interface
type myHandler struct{}

func (myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 解析請求( URL query GET & form POST、PUT)
	// ParseForm() 執行過後 Form 才有資料
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "template.gohtml", data)
}
