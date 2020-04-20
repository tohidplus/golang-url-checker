package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/http/controllers"
	"github.com/tohidplus/url_project/src/app/http/middlewares"
)

type ResultRouter struct {
	controller *controllers.ResultController
}

func (resultRouter ResultRouter) Routes(router *httprouter.Router) {
	router.GET(Prefix("/result/:id"),
		middlewares.Auth(resultRouter.controller.Index),
	)
	router.DELETE(Prefix("/result/:url/:id"),
		middlewares.Auth(
			resultRouter.controller.Destroy,
		),
	)
}
