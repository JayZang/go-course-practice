package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 除了"/"，若都沒有匹配到都會以此 handler 處理
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `<img src="/public/dog.jpg">`)
	})

	// 只要是以 /hello/... 形式都會以這個 handler 處理
	// 為 /hello 也算，預設會導至 /hello/，除非有另外設置 /hello 的 handler
	http.HandleFunc("/hello/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `Hello World`)
	})

	// 只有為 /bye  這個形式才會用此 handler
	// 即使是 /bye/ 也不行
	http.HandleFunc("/bye", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `Bye`)
	})

	// http.FileServer() 的 handler 會直接由設置的根目錄去尋找對應 pattern 的檔案
	// 因此下方代碼若沒有使用 http.StripPrefix 函數的話
	// 系統會去尋找 ./assets/public/dog.jpg，檔案是不存在的
	// 而透過 http.StripPrefix 後會先將 req.URL.PATH 的 '/public' 去除掉 再調用 http.FileServer 的 handler
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
