package errors

import (
	"github.com/tohidplus/url_project/src/helpers/response"
	"net/http"
)

type CredentialError struct{}
type InvalidTokenError struct {
}
type UnauthorizedError struct{}

func (credentialError CredentialError) Print(writer http.ResponseWriter) {
	response.Json(writer, map[string][]string{
		"credentials": {
			"The username or the password is incorrect.",
		},
	}, http.StatusUnprocessableEntity)
}

func (invalidTokenError InvalidTokenError) Print(writer http.ResponseWriter) {
	response.Json(writer, map[string]string{
		"message": "Invalid token.",
	}, http.StatusUnauthorized)
}

func (unauthorizedError UnauthorizedError) Print(writer http.ResponseWriter) {
	response.Json(writer, struct {
		Message string
	}{
		Message: "Unauthorized.",
	}, http.StatusUnauthorized)
}
