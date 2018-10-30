package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// POST and PUT body parameters take precedence over URL query string values.
		q := req.FormValue("q")
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(res, `
			<form action="/" method="POST">
				<input type="text" name="q">
				<input type="submit" value="Submit">
			</form>
			<br>
		`+q)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
