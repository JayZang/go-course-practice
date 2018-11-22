package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserController ...
type UserController struct{}

// NewUserController ...
func NewUserController() *UserController {
	return new(UserController)
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
