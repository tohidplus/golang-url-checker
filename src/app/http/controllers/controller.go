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