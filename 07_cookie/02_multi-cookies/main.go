package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/more", more)

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

func more(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "THIS GENERAL",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "more",
		Value: "THIS MORE",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application/ cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE:", c1)
	}

	c2, err := req.Cookie("general")
	if err == http.ErrNoCookie {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE:", c2)
	}

	c3, err := req.Cookie("more")
	if err == http.ErrNoCookie {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE:", c3)
	}
}
