package exception

import "net/http"

type CustomError interface {
	Print(writer http.ResponseWriter)
}