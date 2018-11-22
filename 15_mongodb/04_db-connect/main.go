package main

import (
	"fmt"
	"net/http"
	"playground/go-web-course/15_mongodb/04_db-connect/controllers"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())
	indexController := controllers.NewIndexController()

	router.GET("/", indexController.Index)

	router.GET("/user/:ID", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user", userController.DeleteUser)

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
