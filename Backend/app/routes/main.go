package routes

import (
	. "api/app/Http/Controllers"
	. "api/app/Http/Middlewares"
	. "api/app/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var env = Lenv.New()

func Router() *mux.Router {
	router := mux.NewRouter()

	pmw := PermissionMiddleware{Token: make(map[string]string)}
	pmw.Populate()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(mux.CORSMethodMiddleware(api))
	api.Use((&HeadersDefaultMiddleware{}).Middleware)

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World! :)"})
	})
	us := api.PathPrefix("/users").Subrouter()
	us.Use(pmw.Middleware)
	us.HandleFunc("", (&UserController{}).Index).Methods("GET")
	us.HandleFunc("", (&UserController{}).Store).Methods("POST")
	us.HandleFunc("/{id}", (&UserController{}).Show).Methods("GET")
	us.HandleFunc("/{id}", (&UserController{}).Update).Methods("PUT")
	us.HandleFunc("/{id}", (&UserController{}).Delete).Methods("DELETE")
	ml := api.PathPrefix("/mail").Subrouter()
	ml.HandleFunc("", (&MailController{}).Store).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("URL_FRONT")},
		AllowCredentials: true,
	})

	handler := c.Handler(api)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))

	return router
}
