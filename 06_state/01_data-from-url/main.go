package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// POST and PUT body parameters take precedence over URL query string values.
		q := req.FormValue("q")
		io.WriteString(res, "The q is "+q)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// visit localhost:8080/?q=hello
