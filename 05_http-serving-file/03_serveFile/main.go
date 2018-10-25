package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `<img src="/dog">`)
	})

	// 返回圖片
	http.HandleFunc("/dog", func(res http.ResponseWriter, req *http.Request) {
		// 和 http.ServeContent() 功能一樣
		// 但簡化了用法
		http.ServeFile(res, req, "dog.jpg")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
