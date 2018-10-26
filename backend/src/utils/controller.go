package utils

import "net/http"

type controllerFunc func(w http.ResponseWriter, r *http.Request)

type Controller struct {
	HandlerFunc controllerFunc
	Path        string
	Method      string
}

func NewController(path string, method string, handlerFunction controllerFunc) *Controller {
	controller := new(Controller)
	controller.Path = "/api/v1/" + path
	controller.Method = method
	controller.HandlerFunc = handlerFunction
	return controller
}
