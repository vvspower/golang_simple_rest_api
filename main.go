package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MustafaAP/ProjectK-backend-Go/router"
)

func main() {
	fmt.Println("ProjectK")

	r := router.Router()
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
	fmt.Println("Listening at port 3000")

}
