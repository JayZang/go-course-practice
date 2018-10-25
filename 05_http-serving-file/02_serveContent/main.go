package main

import (
	"fmt"
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

		fileInfo, err := f.Stat()
		if err != nil {
			http.Error(res, err.Error(), 404)
			return
		}

		// 會使用到 http protocal 緩存技術
		// If modtime is not the zero time or Unix epoch, ServeContent
		// includes it in a `Last-Modified` header in the response. If the
		// request includes an `If-Modified-Since` header, ServeContent uses
		// modtime to decide whether the content needs to be sent at all.
		http.ServeContent(res, req, f.Name(), fileInfo.ModTime(), f)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
