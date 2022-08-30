package controllers

import (
	services "api/app/Services"
	"api/app/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	Message *utils.MessageRequest
}

var service services.UserService

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	resp := service.All()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		uc.Message.WriteMessage(w, "error!", http.StatusNotFound)
		return
	}

	resp, err := service.Get(id)

	if err != nil {
		uc.Message.WriteMessage(w, "not found!", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Store(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	_, err := service.Create(decoder)

	if err != nil {
		uc.Message.WriteMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uc.Message.WriteMessage(w, "successfully created!", http.StatusOK)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	decoder := json.NewDecoder(r.Body)
	resp, err := service.Update(id, decoder)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		uc.Message.WriteMessage(w, "not found", http.StatusNotFound)
		return
	}

	resp, err := service.Delete(id)

	if err != nil || resp == false {
		uc.Message.WriteMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uc.Message.WriteMessage(w, "successfully deleted!", http.StatusOK)
}
