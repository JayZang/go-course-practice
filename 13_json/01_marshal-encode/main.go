package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := `<html>
		<head></head>
		<body>
			At Index
		</body>
	</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		"Chang",
		"Jay",
		[]string{
			"d1",
			"d2",
		},
	}

	json, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		"Li",
		"Han",
		[]string{
			"d1",
			"d2",
		},
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Fatal(err)
	}
}

type person struct {
	FirstName    string
	LastName     string
	Descriptions []string
}
