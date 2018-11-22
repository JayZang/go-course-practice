package main

import (
	"net/http"
	"playground/go-web-course/15_mongodb/03_MVC-struct/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController()
	indexController := controllers.NewIndexController()

	router.GET("/", indexController.Index)

	router.GET("/user/:ID", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user", userController.DeleteUser)

	http.ListenAndServe(":8080", router)
}
