package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/models"
	"github.com/tohidplus/url_project/src/database"
	"github.com/tohidplus/url_project/src/helpers/auth"
	"github.com/tohidplus/url_project/src/helpers/response"
	"net/http"
)

type NotificationController Controller

func (c NotificationController) UrlNotifications(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, url := getAuthUserUrl(request, params.ByName("id"))
	if !ok {
		response.NotFound(writer)
		return
	}
	var notifications [] models.Notification
	database.DB.Model(&url).Related(&models.Notification{}).Find(&notifications)
	response.Json(writer, notifications, http.StatusOK)
}

func (c NotificationController) UserNotifications(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, user := auth.GetUser(request)
	var notifications []models.Notification
	database.DB.Model(&user).Related(&models.Notification{}).Find(&notifications)
	response.Json(writer, notifications, http.StatusOK)
}
