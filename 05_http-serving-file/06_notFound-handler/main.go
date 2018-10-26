package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello World")
	})

	// 瀏覽器會預設請求 /favicon.ico
	// 但我們沒有提供，因此使用 NotFoundHandler() 處理
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
