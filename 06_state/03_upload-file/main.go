package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var s string

		fmt.Println(req.Method)
		if req.Method == http.MethodPost {

			// open received file
			// 解析 http form 傳輸類型 Content-Type: multipart/form-part
			// 若 Content-Type 不為 multipart/form-part 會錯誤
			f, h, err := req.FormFile("q")
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			defer f.Close()

			fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

			// read all file
			bs, err := ioutil.ReadAll(f)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			s = string(bs)

			// create new file
			dst, err := os.Create(filepath.Join("./received-file", h.Filename))
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			// write file
			_, err = dst.Write(bs)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `
			<form action="/" method="POST" enctype="multipart/form-data">
				<input type="file" name="q">
				<input type="submit" value="Submit">
			</form>
			<br>
		`+s)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
