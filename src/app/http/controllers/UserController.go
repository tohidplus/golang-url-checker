package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_project/src/app/models"
	"url_project/src/database"
	"url_project/src/exception"
	"url_project/src/helpers/auth"
	"url_project/src/helpers/auth/errors"
	"url_project/src/helpers/response"
)

type UserController Controller

func (c UserController) Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user := models.ResponseUser{}
	_, authUser := auth.GetUser(request)
	//user:=models.User{}
	database.DB.Where("id = ?", authUser.ID).First(&user)
	if user.ID == 0 {
		response.NotFound(writer)
		return
	}
	response.Json(writer, user, http.StatusOK)
}

func (c UserController) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: models.Password(request.PostFormValue("password")),
	}
	database.DB.Create(&user)
	token, err := auth.GetToken(user)
	exception.LogPrint(err)
	response.Json(writer, token, http.StatusCreated)
}

func (c UserController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ok, token := auth.AttemptLogin(request)
	if !ok {
		errors.CredentialError{}.Print(writer)
		return
	}
	response.Json(writer, token, http.StatusOK)
}
