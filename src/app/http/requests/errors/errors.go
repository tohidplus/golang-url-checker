package errors

import (
	"github.com/tohidplus/url_project/src/helpers/response"
	"net/http"
)

type ValidationErrors map[string][]string

func (vErrors ValidationErrors) Print(writer http.ResponseWriter) {
	response.Json(writer, vErrors,http.StatusUnprocessableEntity)
}
