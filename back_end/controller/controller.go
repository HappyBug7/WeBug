package controller

type Controller struct {
	Shortcut
	Weather
	Schedule
	Search
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
