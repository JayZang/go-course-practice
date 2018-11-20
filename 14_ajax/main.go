package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	indexFile, err := os.Open("./index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	state, err := indexFile.Stat()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fileBs := make([]byte, state.Size())
	_, err = indexFile.Read(fileBs)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(fileBs)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j := `{
		"name": "jay"
	}`

	w.Write([]byte(j))
}
