package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Hello World^*&)($%#"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
