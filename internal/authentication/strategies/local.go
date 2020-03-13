package authentication

import (
	"awesome-go/internal/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterLocalStrategy represent initial functiona when application is created
func RegisterLocalStrategy(router *mux.Router) {
	router.HandleFunc("/sign-in", SignIn).Methods(http.MethodGet)
}

// SignIn will response contacts as an array
func SignIn(w http.ResponseWriter, r *http.Request) {
	contact := []models.Contact{
		{
			FirstName: "Jugkapongsss",
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
