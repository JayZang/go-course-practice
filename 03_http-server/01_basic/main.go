package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := myHandler{}

	// 建立 tcp server，應用層為 http
	http.ListenAndServe(":8080", handler)
}

// 實作 Handler Interface
type myHandler struct{}

func (myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
