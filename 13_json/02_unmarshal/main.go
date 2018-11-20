package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := `{
		"OtherField": "Hey",
		"FirstName": "Chang",
		"LastName": "Jay",
		"Description": [
			"p1",
			"p2"
		]
	}`

	var data person
	if err := json.Unmarshal([]byte(j), &data); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	j = `[{
		"OtherField": "Hey",
		"FirstName": "Chang",
		"LastName": "Jay",
		"Description": [
			"p1",
			"p2"
		]
	}, {
		"OtherField": "OKOK",
		"FirstName": "Li",
		"LastName": "Han",
		"Description": [
			"p1",
			"p2"
		]
	}]`

	var people []person
	if err := json.Unmarshal([]byte(j), &people); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(people)
	}
}

type person struct {
	FirstName   string
	LastName    string
	Description []string
}
