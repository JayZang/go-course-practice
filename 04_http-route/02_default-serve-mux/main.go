package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 使用預設的 Serve Mux
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello World")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
