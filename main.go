package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/MustafaAP/ProjectK-backend-Go/router"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	fmt.Println("ProjectK")

	r := router.Router()
	handler := c.Handler(r)

	// fmt.Sprintf(":%s", os.Getenv("PORT")) use this when deploy
	//Socket.Socket()
	log.Fatal(http.ListenAndServe(":9000", handler))
	fmt.Println("Listening at port 3000")

}
