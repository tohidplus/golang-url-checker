package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/helpers/response"
	"net/http"
)

type IndexController Controller

func (c IndexController) Index(writer http.ResponseWriter,request *http.Request,params httprouter.Params) {
	response.Json(writer, map[string][]string{
		"email":{
			"asdasdasd",
		},
	},http.StatusOK)
}