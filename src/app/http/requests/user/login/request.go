package login

import (
	"github.com/julienschmidt/httprouter"
	errors2 "github.com/tohidplus/url_project/src/app/http/requests/errors"
	"github.com/tohidplus/url_project/src/app/http/rules"
	"net/http"
)

func UserLoginRequest(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		errors := errors2.ValidationErrors{}
		checkEmail(request, errors)
		checkPassword(request, errors)
		if len(errors) > 0 {
			errors.Print(writer)
			return
		}
		next(writer, request, params)
	}
}

func checkPassword(request *http.Request, errors errors2.ValidationErrors) {
	password := request.PostFormValue("password")
	if password == "" {
		errors["password"] = append(errors["password"], "The password field is required.")
	}
}

func checkEmail(request *http.Request, errors errors2.ValidationErrors) {
	email := request.PostFormValue("email")
	if email == "" {
		errors["email"] = append(errors["email"], "The email field is required.")
	}
	if email != "" && !rules.CheckEmail(email) {
		errors["email"] = append(errors["email"], "The email field is not a valid email type.")
	}
}
