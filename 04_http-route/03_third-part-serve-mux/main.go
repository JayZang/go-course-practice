package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	mux := httprouter.New()

	// RESTful route
	mux.GET("/", index)
	mux.GET("/user/:name", user)
	mux.POST("/user", createUser)
	mux.GET("/signup", signUP)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(res, "Hello World!")
}

func user(res http.ResponseWriter, req *http.Request, para httprouter.Params) {
	fmt.Fprintln(res, "Hello, "+para.ByName("name"))
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm()
	fmt.Fprintln(res, "User "+req.PostForm.Get("name")+" created")
}

func signUP(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "template.gohtml", nil)
	if err != nil {
		panic(err)
	}
}
