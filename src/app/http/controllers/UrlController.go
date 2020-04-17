package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"url_project/src/app/models"
	"url_project/src/database"
	"url_project/src/helpers/auth"
	"url_project/src/helpers/response"
)

type UrlController Controller

func (c UrlController) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var urls []models.Url
	_, user := auth.GetUser(request)
	database.DB.Where("user_id = ?", user.ID).Find(&urls)
	response.Json(writer, urls, http.StatusOK)
}

func (c UrlController) Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fields := getParsedFields(request)
	url := models.Url{
		UserID:    fields.UserID,
		Path:      request.PostFormValue("path"),
		Method:    request.PostFormValue("method"),
		Headers:   fields.Headers,
		Body:      fields.Body,
		Threshold: fields.Threshold,
	}
	database.DB.Create(&url)
	response.Json(writer, url, http.StatusCreated)
}

func (c UrlController) Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params)
	if !ok {
		response.NotFound(writer)
		return
	}
	response.Json(writer, url, http.StatusOK)
}

func (c UrlController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params)
	if !ok {
		response.NotFound(writer)
		return
	}
	fields := getParsedFields(request)
	url.Path = request.PostFormValue("path")
	url.Headers = fields.Headers
	url.Body = fields.Body
	url.Threshold = fields.Threshold
	database.DB.Save(&url)
	response.Json(writer, url, http.StatusOK)
}

func (c UrlController) Destroy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params)
	if !ok {
		response.NotFound(writer)
		return
	}
	database.DB.Delete(url)
	response.Json(writer, "", http.StatusNoContent)
}

type parsedFields struct {
	UserID    uint
	Headers   models.JsonProperty
	Body      models.JsonProperty
	Threshold uint
}

func getParsedFields(request *http.Request) parsedFields {
	_, user := auth.GetUser(request)
	headers := models.JsonProperty{}
	body := models.JsonProperty{}
	_ = json.Unmarshal([]byte(request.PostFormValue("headers")), &headers)
	_ = json.Unmarshal([]byte(request.PostFormValue("body")), &body)
	threshold, _ := strconv.ParseUint(request.PostFormValue("threshold"), 10, 64)
	return parsedFields{
		UserID:    user.ID,
		Headers:   headers,
		Body:      body,
		Threshold: uint(threshold),
	}
}

func getAuthUserUrl(request *http.Request, params httprouter.Params) (bool, models.Url) {
	id := params.ByName("id")
	_, user := auth.GetUser(request)
	url := models.Url{}
	database.DB.Where(map[string]interface{}{"id": id, "user_id": user.ID}).First(&url)
	if url.ID == 0 {
		return false, url
	}
	return true, url
}
