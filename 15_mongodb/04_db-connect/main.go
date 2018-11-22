package main

import (
	"fmt"
	"html/template"
	"net/http"
	"playground/go-web-course/15_mongodb/04_db-connect/controllers"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func main() {
	tpl := template.Must(template.ParseGlob("./templates/*"))
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())
	indexController := controllers.NewIndexController(tpl)

	router.GET("/", indexController.Index)

	router.POST("/user", userController.CreateUser)
	router.GET("/user/:ID", userController.GetUser)
	router.DELETE("/user/:ID", userController.DeleteUser)

	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connect success")
	return s
}
