package main

import (
	"awesome-go/internal/authentication"
	"awesome-go/internal/configs"
	"awesome-go/internal/controllers"
	logger "awesome-go/internal/pkg"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// HTTP will start http server
func HTTP(port string) error {
	router := mux.NewRouter()

	if err := controllers.Register(router); err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	if err := authentication.Register(router); err != nil {
		log.Fatalf("error initializing auth strategies: %v", err)
	}

	return http.ListenAndServe(port, router)
}

func main() {
	configs.Init()
	logger.Init()

	log.Info("Starting go application on http://localhost:", viper.Get("port"))
	if err := HTTP; err != nil {
		port := fmt.Sprintf(":%s", viper.Get("port"))
		log.Fatal(err(port))
	}
}
