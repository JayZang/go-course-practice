package controllers

import (
	"encoding/json"
	"net/http"
	"playground/go-web-course/15_mongodb/04_db-connect/models"

	"gopkg.in/mgo.v2/bson"

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
	id := param.ByName("ID")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("go-web-course").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jn, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jn)
}

// CreateUser ...
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	u := models.User{}

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return
	}

	// 給予使用者唯一ID
	u.ID = bson.NewObjectId()

	// 新增使用者至DB
	uc.session.DB("go-web-course").C("users").Insert(u)

	jn, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jn)
}

// DeleteUser ...
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	id := param.ByName("ID")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("go-web-course").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
