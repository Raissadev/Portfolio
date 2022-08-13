package routes

import (
	utils "api/app/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var env utils.LoadEnv

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api/v1")

	// router.HandleFunc("/", YourHandler).Methods("GET")

	env.New()

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))

	return router
}
