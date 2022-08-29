package controllers

import (
	services "api/app/Services"
	"encoding/json"
	"net/http"
)

type UserController struct {
}

var service services.UserService

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := service.All()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	//
}
