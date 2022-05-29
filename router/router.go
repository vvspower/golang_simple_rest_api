package router

import (
	"github.com/MustafaAP/ProjectK/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")

	return router
}
