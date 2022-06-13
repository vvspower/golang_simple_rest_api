package router

import (
	"github.com/MustafaAP/ProjectK-backend-Go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	//user
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUserWithAuth).Methods("GET")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/init-igd", controllers.InitGameDetails).Methods("POST")
	router.HandleFunc("/update-igd", controllers.UpdateIGD).Methods("PUT")

	// friends

	router.HandleFunc("/sendfq/{id}", controllers.FriendReq).Methods("POST")
	router.HandleFunc("/getfq", controllers.GetFriendReqs).Methods("GET")
	router.HandleFunc("/acceptfq", controllers.AcceptFriendReq).Methods("POST")
	router.HandleFunc("/deletefq/{id}", controllers.DeleteFriend).Methods("DELETE")
	router.HandleFunc("/getfriends", controllers.GetFriends).Methods("GET")
	router.HandleFunc("/friend", controllers.GetFriend).Methods("POST")

	// listing

	router.HandleFunc("/listing", controllers.CreateListing).Methods("POST")
	router.HandleFunc("/listing/{user}", controllers.GetUserListings).Methods("GET")
	router.HandleFunc("/listing", controllers.GetAllListing).Methods("GET")
	router.HandleFunc("/onelist/{id}", controllers.GetOneListing).Methods("GET")

	// update listing
	router.HandleFunc("/listing/{id}", controllers.UpdateListing).Methods("PUT")

	// delete listing
	router.HandleFunc("/listing/{id}", controllers.DeleteListing).Methods("DELETE")

	// socket

	router.HandleFunc("/api/messages/{channel}", controllers.WsConn).Methods("POST")

	// valorant api

	router.HandleFunc("/api/valorant/{region}/{user}/{tag}", controllers.GetMMRData).Methods("GET")

	return router
}
