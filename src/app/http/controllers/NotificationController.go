package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_project/src/app/models"
	"url_project/src/database"
	"url_project/src/helpers/auth"
	"url_project/src/helpers/response"
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
