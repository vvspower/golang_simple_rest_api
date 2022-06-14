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
		AppID:   "-",
		Key:     "-",
		Secret:  "-",
		Cluster: "-",
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
