package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexController ...
type IndexController struct {
	tpl *template.Template
}

// NewIndexController ...
func NewIndexController(tpl *template.Template) *IndexController {
	return &IndexController{tpl}
}

// Index ...
func (ic IndexController) Index(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	ic.tpl.ExecuteTemplate(w, "index.ghtml", nil)
}
