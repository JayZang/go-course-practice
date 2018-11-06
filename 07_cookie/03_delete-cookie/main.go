package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1>
			<a href="/set">Set Cookie</a>
		</h1>`)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "stuff",
	})

	fmt.Fprintln(res, `<h1>
			<a href="/read">Read Cookie</a>
		</h1>`)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(res, `<h1>
			<a href="/expire">Expire Cookie</a>
		</h1>`)
	fmt.Fprintln(res, c)
}

func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	c.MaxAge = -1
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}
