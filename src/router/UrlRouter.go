package router

import (
	"github.com/julienschmidt/httprouter"
	"url_project/src/app/http/controllers"
	"url_project/src/app/http/middlewares"
	"url_project/src/app/http/requests/url/store"
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
	router.DELETE(Prefix("/url/:id"),
		middlewares.Auth(
			urlRouter.controller.Destroy,
		),
	)
}
