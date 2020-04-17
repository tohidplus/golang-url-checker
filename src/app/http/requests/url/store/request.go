package store

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	errors2 "url_project/src/app/http/requests/errors"
	"url_project/src/app/http/rules"
)

func UrlStoreRequest(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		errors := errors2.ValidationErrors{}
		checkPath(request, errors)
		checkMethod(request, errors)
		checkHeaders(request, errors)
		checkBody(request, errors)
		checkThreshold(request, errors)
		if len(errors) > 0 {
			errors.Print(writer)
			return
		}
		next(writer, request, params)
	}
}

func checkThreshold(request *http.Request, errors errors2.ValidationErrors) {
	threshold := request.PostFormValue("threshold")
	if threshold == "" {
		errors["threshold"] = append(errors["threshold"], "The threshold field is required.")
	}
	if threshold != "" && !rules.CheckIsAnInteger(threshold) {
		errors["threshold"] = append(errors["threshold"], "The threshold field must be an integer value.")
	}
}

func checkBody(request *http.Request, errors errors2.ValidationErrors) {
	body := request.PostFormValue("body")
	if body == "" {
		errors["body"] = append(errors["body"], "The body field is required.")
	}
	if body != "" && !rules.CheckJsonString(body) {
		errors["body"] = append(errors["body"], "The body field is must be a valid json string.")
	}
}

func checkHeaders(request *http.Request, errors errors2.ValidationErrors) {
	headers := request.PostFormValue("headers")
	if headers == "" {
		errors["headers"] = append(errors["headers"], "The headers field is required.")
	}
	if headers != "" && !rules.CheckJsonString(headers) {
		errors["headers"] = append(errors["headers"], "The headers field is must be a valid json string.")
	}
}

func checkMethod(request *http.Request, errors errors2.ValidationErrors) {
	method := request.PostFormValue("method")
	if method == "" {
		errors["method"] = append(errors["method"], "The method field is required.")
	}
	if method != "" && !rules.CheckInArrayString(method, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}) {
		errors["method"] = append(errors["method"], "The method field is not valid.")
	}
}

func checkPath(request *http.Request, errors errors2.ValidationErrors) {
	path := request.PostFormValue("path")
	if path == "" {
		errors["path"] = append(errors["path"], "The path field is required.")
	}
	if path != "" && !rules.CheckUrl(path) {
		errors["path"] = append(errors["path"], "The path field is not a valid url.")
	}
}
