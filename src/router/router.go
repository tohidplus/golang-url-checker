package router

import (
	"github.com/julienschmidt/httprouter"
	"url_project/src/app/http/controllers"
)

type HttpRoutes interface {
	Routes(router *httprouter.Router)
}

func newIndexRouter() IndexRouter {
	return IndexRouter{
		controllers.NewIndexController(),
	}
}

func newUserRouter() UserRouter {
	return UserRouter{
		controllers.NewUserController(),
	}
}

func newUrlRouter() UrlRouter {
	return UrlRouter{
		controller: controllers.NewUrlController(),
	}
}

func Prefix(path string) string {
	return "/api" + path
}

func RegisterHttpRoutes() *httprouter.Router {

	r := httprouter.New()
	newIndexRouter().Routes(r)
	newUserRouter().Routes(r)
	newUrlRouter().Routes(r)
	return r
}
