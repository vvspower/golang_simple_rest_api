package controllers

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

//
//func wsConn(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	at := r.Header.Get("auth-token")
//	data, _ := helper.ExtractClaims(at)
//	userid := data["userid"] // this is the id of client
//	userid = fmt.Sprintf(":%s", userid)
//
//	params := mux.Vars(r)
//	id := params["id"] // this is id of person who the client is trying to chat to
//
//	//	TODO find the channel created between those users and establish connection by sending back the channel and ap
//
//}
