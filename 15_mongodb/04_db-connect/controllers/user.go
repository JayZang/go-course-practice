package controllers

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

// UserController ...
type UserController struct {
	session *mgo.Session
}

// NewUserController ...
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser ...
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}

// CreateUser ...
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}

// DeleteUser ...
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}
