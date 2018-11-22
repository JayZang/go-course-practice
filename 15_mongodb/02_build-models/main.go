package main

import (
	"encoding/json"
	"net/http"

	"playground/go-web-course/15_mongodb/02_build-models/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/user/:ID", getUser)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	html := `
		<html>
			<head></head>
			<body>
				<a href="/user/89ysd9qdjs">
					User Link
				</a>
			</body>
		</html>
	`
	w.Write([]byte(html))
}

func getUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	u := models.User{
		Name:   "Jay",
		Gender: "male",
		Age:    22,
		ID:     param.ByName("ID"),
	}

	j, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
