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
	http.Redirect(res, req, "/success", http.StatusTemporaryRedirect)
}

func success(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Success")
}
