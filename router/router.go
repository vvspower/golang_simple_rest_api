package router

import (
	"github.com/MustafaAP/ProjectK-backend-Go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")

	// listing

	router.HandleFunc("/listing", controllers.CreateListing).Methods("POST")
	router.HandleFunc("/listing/{user}", controllers.GetUserListings).Methods("GET")
	router.HandleFunc("/listing", controllers.GetAllListing).Methods("GET")
	router.HandleFunc("/onelist/{id}", controllers.GetOneListing).Methods("GET")

	return router
}
