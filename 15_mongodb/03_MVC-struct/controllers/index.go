package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexController ...
type IndexController struct{}

// NewIndexController ...
func NewIndexController() *IndexController {
	return new(IndexController)
}

// Index ...
func (ic IndexController) Index(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}
