package controllers

import (
	"awesome-go/internal/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterContact represent initial functiona when application is created
func RegisterContact(router *mux.Router) {
	router.HandleFunc("/contacts", List).Methods(http.MethodGet)
}

// List will response contacts as an array
func List(w http.ResponseWriter, r *http.Request) {
	contact := []models.Contact{
		{
			FirstName: "Jugkapong",
			LastName:  "Pooban",
			Email:     []string{"j.pooban@gmail.com", "pise_acsp_1@hotmail.com"},
		},
		{
			FirstName: "Jirapat",
			LastName:  "Charoenchai",
			Email:     []string{"pimpat.1705@gmail.com", "Pim_jirapat1705@hotmail.com"},
		},
	}

	js, err := json.Marshal(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
