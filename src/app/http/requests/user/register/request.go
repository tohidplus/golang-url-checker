package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	errors2 "url_project/src/app/http/requests/errors"
	"url_project/src/app/http/rules"
)

func UserRegisterRequest(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var errors = errors2.ValidationErrors{}
		checkEmail(request, errors)
		checkName(request, errors)
		checkPassword(request, errors)
		if len(errors) > 0 {
			errors.Print(writer)
			return
		}
		next(writer, request, params)
	}
}

func checkPassword(request *http.Request, errors errors2.ValidationErrors) {
	if request.PostFormValue("password") == "" {
		errors["password"] = append(errors["password"], "The password field is required.")
	}
	if request.PostFormValue("password") != "" &&
		len(request.PostFormValue("password")) < 6 {
		errors["password"] = append(errors["password"], "The password field must be at least 6 characters.")
	}
	if request.PostFormValue("password") != "" &&
		request.PostFormValue("password") != request.PostFormValue("password_confirmation") {
		errors["password"] = append(errors["password"], "The password filed does not match the confirmation value.")
	}
}

func checkName(request *http.Request, errors errors2.ValidationErrors) {
	if request.PostFormValue("name") == "" {
		errors["name"] = append(errors["name"], "The name field is required.")
	}
}

func checkEmail(request *http.Request, errors errors2.ValidationErrors) {
	if request.PostFormValue("email") != "" && !rules.CheckEmail(request.PostFormValue("email")) {
		errors["email"] = append(errors["email"], "The email field is not a valid email type.")
	}
	if request.PostFormValue("email") == "" {
		errors["email"] = append(errors["email"], "The email field is required.")
	}
	if rules.CheckEmail(request.PostFormValue("email")) &&
		!rules.CheckUniqueness("users", "email", request.PostFormValue("email")) {
		errors["email"] = append(errors["email"], "The email field has already been taken.")
	}
}

