package main

import (
	"awesome-go/internal/authentication"
	"awesome-go/internal/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTP will start http server
func HTTP() error {
	router := mux.NewRouter()

	if err := controllers.Register(router); err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	if err := authentication.Register(router); err != nil {
		log.Fatalf("error initializing auth strategies: %v", err)
	}

	return http.ListenAndServe(":8080", router)
}

func main() {
	log.Println("Go application is starting on http://localhost:8080")

	if err := HTTP; err != nil {
		log.Fatal(err())
	}
}
