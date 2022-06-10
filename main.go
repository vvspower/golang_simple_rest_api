package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MustafaAP/ProjectK-backend-Go/router"
)

func main() {
	fmt.Println("ProjectK")

	r := router.Router()

	// fmt.Sprintf(":%s", os.Getenv("PORT")) use this when deploy
	//Socket.Socket()
	log.Fatal(http.ListenAndServe(":9000", r))
	fmt.Println("Listening at port 3000")

}
