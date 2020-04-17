package controllers

type Controller struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
func NewUserController() *UserController {
	return &UserController{}
}
func NewUrlController() *UrlController {
	return &UrlController{}
}
func NewResultController() *ResultController {
	return &ResultController{}
}
func NewNotificationController() *NotificationController {
	return &NotificationController{}
}
