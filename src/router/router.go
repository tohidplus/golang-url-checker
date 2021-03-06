package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/http/controllers"
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

func newResultRouter() ResultRouter {
	return ResultRouter{
		controller: controllers.NewResultController(),
	}
}

func newNotificationRouter() NotificationRouter {
	return NotificationRouter{controller: controllers.NewNotificationController()}
}

func Prefix(path string) string {
	return "/api" + path
}

func RegisterHttpRoutes() *httprouter.Router {

	r := httprouter.New()
	newIndexRouter().Routes(r)
	newUserRouter().Routes(r)
	newUrlRouter().Routes(r)
	newResultRouter().Routes(r)
	newNotificationRouter().Routes(r)
	return r
}
