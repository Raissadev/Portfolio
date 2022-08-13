package controllers

import (
	services "api/app/Services"
	"net/http"
)

type UserController struct {
}

var service services.UserService

func (uc *UserController) Index() string {
	return service.All()

}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	//
}
