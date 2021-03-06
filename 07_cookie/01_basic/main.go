package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)

	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "stuff",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application/ cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(res, "YOUR COOKIE:", c)
}
