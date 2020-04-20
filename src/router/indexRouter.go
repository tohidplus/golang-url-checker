package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/http/controllers"
)

type IndexRouter struct {
	controller *controllers.IndexController
}


func (indexRouter IndexRouter) Routes(router *httprouter.Router) {
	router.GET(Prefix("/"), indexRouter.controller.Index)
}
