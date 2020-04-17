package router

import (
	"github.com/julienschmidt/httprouter"
	"url_project/src/app/http/controllers"
	"url_project/src/app/http/middlewares"
	"url_project/src/app/http/requests/user/login"
	"url_project/src/app/http/requests/user/register"
)

type UserRouter struct {
	controller *controllers.UserController
}

func (userRouter UserRouter) Routes(router *httprouter.Router){
	router.GET(Prefix("/user"),middlewares.Auth(userRouter.controller.Show))
	router.POST(Prefix("/user/register"), register.UserRegisterRequest(userRouter.controller.Register))
	router.POST(Prefix("/user/login"),login.UserLoginRequest(userRouter.controller.Login))
}