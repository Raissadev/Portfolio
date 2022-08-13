package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api/v1")

	// router.HandleFunc("/", YourHandler).Methods("GET")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))

	return router
}
