package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_project/src/helpers/auth"
	"url_project/src/helpers/auth/errors"
)

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ok, _ := auth.GetUser(request)
		if !ok {
			errors.UnauthorizedError{}.Print(writer)
			return
		}
		next(writer, request, params)
	}
}
