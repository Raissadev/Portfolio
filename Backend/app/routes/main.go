package routes

import (
	. "api/app/Http/Controllers"
	. "api/app/Http/Middlewares"
	. "api/app/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var env = Lenv

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
	us.HandleFunc("", (&UserController{}).Index).Methods("GET")
	us.HandleFunc("", (&UserController{}).Store).Methods("POST")
	us.HandleFunc("/{id}", (&UserController{}).Show).Methods("GET")
	us.HandleFunc("/{id}", (&UserController{}).Update).Methods("PUT")
	us.HandleFunc("/{id}", (&UserController{}).Delete).Methods("DELETE")
	ml := r.PathPrefix("/mail").Subrouter()
	ml.HandleFunc("", (&MailController{}).Store).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("URL_FRONT")},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), handler))

	return router
}
