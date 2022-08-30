package routes

import (
	controllers "api/app/Http/Controllers"
	utils "api/app/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var env utils.LoadEnv
var controller controllers.UserController

func Router() *mux.Router {
	router := mux.NewRouter()
	// router.PathPrefix("/api/v1")

	router.HandleFunc("/", controller.Index).Methods("GET")
	router.HandleFunc("/", controller.Store).Methods("POST")
	router.HandleFunc("/{id}", controller.Show).Methods("GET", "PUT", "DELETE")

	env.New()

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))

	return router
}
