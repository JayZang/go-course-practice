package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World!\n"))
}
