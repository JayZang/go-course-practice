package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/process", process)
	http.HandleFunc("/success", success)

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	html := `
		<html>
			<head></head>
			<body>
				<form action="/process" method="POST">
					<input type="text" name="val">
					<input type="submit">
				</form>
			</body>
		</html>`

	fmt.Fprintln(res, html)
}

func process(res http.ResponseWriter, req *http.Request) {
	// 設置 Location header，並返回指定之 Status Code
	// 307 status：瀏覽器看到此URL還是會造訪，可當作特定處理程序後導向其他位址
	// 301 status：瀏覽器看到此URL後會自動記住目標位址，下次不再訪問此URL
	res.Header().Set("Location", "/success")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func success(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Success")
}
