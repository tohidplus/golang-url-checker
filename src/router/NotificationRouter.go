package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tohidplus/url_project/src/app/http/controllers"
	"github.com/tohidplus/url_project/src/app/http/middlewares"
)

type NotificationRouter struct {
	controller *controllers.NotificationController
}

func (notificationRouter NotificationRouter) Routes(router *httprouter.Router) {
	router.GET(Prefix("/notification/url/:id"),
		middlewares.Auth(
			notificationRouter.controller.UrlNotifications,
		),
	)
	router.GET(Prefix("/notification/user/:id"),
		middlewares.Auth(
			notificationRouter.controller.UserNotifications,
		),
	)
}
