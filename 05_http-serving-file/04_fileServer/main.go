package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Pic", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `<img src="/dog.jpg">`)
	})

	// 可顯示目錄檔案架構...
	// 在所設置的檔案系統根目錄放置 index.html 以防止顯示檔案架構
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
