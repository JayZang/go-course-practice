package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := myHandler{}

	// 自訂新增的 Serve Mux
	serveMux := http.NewServeMux()
	serveMux.Handle("/", handler)

	err := http.ListenAndServe(":8080", serveMux)
	if err != nil {
		panic(err)
	}
}

type myHandler struct{}

func (myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}
