package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pusher/pusher-http-go"
	"net/http"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func WsConn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	channel := params["channel"]
	fmt.Println(channel)

	pusherClient := pusher.Client{
		AppID:   "1421095",
		Key:     "3a4c1b62e8e7b86334fd",
		Secret:  "d445191d82cd77c696de",
		Cluster: "ap2",
		Secure:  true,
	}

	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}

	pusherClient.Trigger(channel, "message", data)

	json.NewEncoder(w).Encode("Connected")

	//	TODO find the channel created between those users and establish connection by sending back the channel and ap
}
