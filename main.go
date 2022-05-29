package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MustafaAP/ProjectK/router"
)

func main() {
	fmt.Println("ProjectK")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":9000", r))
	fmt.Println("Listening at port 3000")

}
