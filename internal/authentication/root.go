package authentication

import (
	authentication "awesome-go/internal/authentication/strategies"

	"github.com/gorilla/mux"
)

// Register will assign all of auth strategies
func Register(router *mux.Router) error {
	auth := router.PathPrefix("/auth").Subrouter()
	authentication.RegisterLocalStrategy(auth)

	return nil
}
