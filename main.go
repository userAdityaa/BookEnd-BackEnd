package main

import (
	"log"
	"net/http"

	"github.com/userAdityaa/BookEndBackEnd/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Server is running in port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
