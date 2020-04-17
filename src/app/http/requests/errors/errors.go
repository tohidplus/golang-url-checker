package errors

import (
	"net/http"
	"url_project/src/helpers/response"
)

type ValidationErrors map[string][]string

func (vErrors ValidationErrors) Print(writer http.ResponseWriter) {
	response.Json(writer, vErrors,http.StatusUnprocessableEntity)
}
