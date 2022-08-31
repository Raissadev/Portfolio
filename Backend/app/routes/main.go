package routes

import (
	. "api/app/Http/Controllers"
	. "api/app/Http/Middlewares"
	utils "api/app/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var env utils.LoadEnv
var userController UserController
var mailController MailController

func Router() *mux.Router {
	pmw := PermissionMiddleware{Token: make(map[string]string)}
	pmw.Populate()
	headers := HeadersDefaultMiddleware{}

	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(headers.Middleware)

	r := router.PathPrefix("/api/v1").Subrouter()

	us := r.PathPrefix("/users").Subrouter()
	us.Use(pmw.Middleware)
	us.HandleFunc("", userController.Index).Methods("GET")
	us.HandleFunc("", userController.Store).Methods("POST")
	us.HandleFunc("/{id}", userController.Show).Methods("GET")
	us.HandleFunc("/{id}", userController.Update).Methods("PUT")
	us.HandleFunc("/{id}", userController.Delete).Methods("DELETE")

	ml := r.PathPrefix("/mail").Subrouter()
	ml.HandleFunc("", mailController.Store).Methods("POST")

	env.New()
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))

	return router
}
