package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/http/controllers"
	"github.com/tohidplus/url_project/src/app/http/middlewares"
	"github.com/tohidplus/url_project/src/app/http/requests/url/store"
)

type UrlRouter struct {
	controller *controllers.UrlController
}

/*
|-------------------------------------------------------
| Url routes
|-------------------------------------------------------
*/
func (urlRouter UrlRouter) Routes(router *httprouter.Router) {
	router.GET(Prefix("/url"),
		middlewares.Auth(
			urlRouter.controller.Index,
		),
	)
	router.POST(Prefix("/url"),
		middlewares.Auth(
			store.UrlStoreRequest(
				urlRouter.controller.Store,
			),
		),
	)
	router.GET(Prefix("/url/:id"),
		middlewares.Auth(
			urlRouter.controller.Show,
		),
	)
	router.PATCH(Prefix("/url/:id"),
		middlewares.Auth(
			store.UrlStoreRequest(
				urlRouter.controller.Update,
			),
		),
	)
	router.GET(Prefix("/url/:id/call"),
		middlewares.Auth(
			urlRouter.controller.Call,
		),
	)
	router.DELETE(Prefix("/url/:id"),
		middlewares.Auth(
			urlRouter.controller.Destroy,
		),
	)
}
