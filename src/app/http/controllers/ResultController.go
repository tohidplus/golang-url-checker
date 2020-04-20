package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/models"
	"github.com/tohidplus/url_project/src/database"
	"github.com/tohidplus/url_project/src/helpers/response"
	"net/http"
)

type ResultController Controller

func (c ResultController) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params.ByName("id"))
	if !ok {
		response.NotFound(writer)
		return
	}
	var results []models.Result
	database.DB.Where("url_id = ?", url.ID).Find(&results)
	response.Json(writer, results, http.StatusOK)
}

func (c ResultController) Destroy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params.ByName("url"))
	id := params.ByName("id")
	if !ok {
		response.NotFound(writer)
		return
	}
	var result models.Result
	database.DB.Where(map[string]interface{}{"id": id, "url_id": url.ID}).First(&result)
	if result.ID == 0 {
		response.NotFound(writer)
		return
	}
	database.DB.Delete(&result)
	response.Json(writer, "", http.StatusNoContent)
}
