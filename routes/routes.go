package routes

import (
	ControllerUser "api/controller"
	MiddlewareAuth "api/middleware"

	"github.com/gorilla/mux"
)

// Init - inicializa as rotas
func Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", ControllerUser.Create).Methods("POST")
	router.HandleFunc("/login", ControllerUser.Login).Methods("POST")

	// Rotas Autenticadas
	router.HandleFunc("/user", MiddlewareAuth.Auth(ControllerUser.Update)).Methods("PUT")
	router.HandleFunc("/user", MiddlewareAuth.Auth(ControllerUser.List)).Methods("GET")
	router.HandleFunc("/user/{id}", MiddlewareAuth.Auth(ControllerUser.Find)).Methods("GET")
	router.HandleFunc("/user", MiddlewareAuth.Auth(ControllerUser.Delete)).Methods("DELETE")
	return router
}
