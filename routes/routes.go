package routes

import (
	ControllerUser "api/controller"
	MiddlewareAuth "api/middleware"

	"github.com/gorilla/mux"
)

// Init - inicializa as rotas
func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", ControllerUser.Login).Methods("POST")
	router.HandleFunc("/user", ControllerUser.List).Methods("GET")
	router.HandleFunc("/user/{id}", MiddlewareAuth.Auth(ControllerUser.Find)).Methods("GET")
	router.HandleFunc("/user/{id}", ControllerUser.Create).Methods("POST")
	router.HandleFunc("/user/{id}", ControllerUser.Delete).Methods("DELETE")
	return router
}
