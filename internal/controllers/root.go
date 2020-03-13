package controllers

import (
	"awesome-go/internal/controllers/v1"

	"github.com/gorilla/mux"
)

// Register will assign all of route in application
func Register(router *mux.Router) error {
	api := router.PathPrefix("/api").Subrouter()

	// List of api version 1
	v1 := api.PathPrefix("/v1").Subrouter()
	controllers.RegisterContact(v1)

	return nil
}
