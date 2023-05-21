package routes

import (
	"log"
	"net/http"

	"github.com/aleroxac/alura-golang/controllers"
	"github.com/aleroxac/alura-golang/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/api/v1/healthz", controllers.Healthcheck).Methods("GET")
	router.HandleFunc("/api/v1/skills", controllers.List).Methods("GET")
	router.HandleFunc("/api/v1/skills/{name}", controllers.GetByName).Methods("GET")
	router.HandleFunc("/api/v1/skills", controllers.Create).Methods("POST")
	router.HandleFunc("/api/v1/skills/{name}", controllers.Update).Methods("PUT")
	router.HandleFunc("/api/v1/skills/{name}", controllers.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
