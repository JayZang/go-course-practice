package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

type User struct {
	Name string
	Age  string
}

var sessionDB = map[string]string{}
var userDB = map[string]User{}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/user", user)

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var userName = req.FormValue("name")
		var userAge = req.FormValue("age")
		var user = User{
			userName,
			userAge,
		}

		id, _ := uuid.NewV4()
		userDB[userName] = user
		sessionDB[id.String()] = userName

		cookie := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(res, cookie)
		http.Redirect(res, req, "/user", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintln(res, `
		<html>
			<head></head>
			<body>
				<form action="/" method="POST">
					<input name="name" placeholder="Name">
					<input name="age" placeholder="Age">
					<input type="submit">
				</form>
			</body>
		</html>
	`)
}

func user(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	uuid := cookie.Value
	userName := sessionDB[uuid]
	user := userDB[userName]

	fmt.Fprintln(res, user)
}
