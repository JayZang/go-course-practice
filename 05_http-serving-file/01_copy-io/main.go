package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `<img src="/dog">`)
	})

	// 返回圖片
	http.HandleFunc("/dog", func(res http.ResponseWriter, req *http.Request) {
		f, err := os.Open("./dog.jpg")
		if err != nil {
			http.Error(res, err.Error(), 404)
			return
		}
		defer f.Close()

		io.Copy(res, f)
	})

	http.ListenAndServe(":8080", nil)
}
