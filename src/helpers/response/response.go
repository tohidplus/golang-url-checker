package response

import (
	"encoding/json"
	"net/http"
	"url_project/src/exception"
)

func Json(writer http.ResponseWriter, data interface{}, statusCode int) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(statusCode)
	jData, err := json.Marshal(data)
	exception.LogPrint(err)
	_, err = writer.Write(jData)
	exception.LogPrint(err)
}

func NotFound(writer http.ResponseWriter) {
	Json(writer, struct {
		Message string `json:"message"`
	}{
		Message: "Not found.",
	}, http.StatusNotFound)
}
