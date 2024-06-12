package authenticationRoutes

import (
	"github.com/ahsansandiah/dealls-test/packages/manager"
	"github.com/gorilla/mux"

	authHandler "github.com/ahsansandiah/dealls-test/handlers/authentication/delivery/handler"
)

func NewRoutes(r *mux.Router, mgr manager.Manager) {
	route := r.PathPrefix("").Subrouter()
	route.Use(mgr.GetMiddleware().CheckToken)

	authHandler := authHandler.NewAuthHandler(mgr)
	r.Handle("/sign-up", authHandler.SignUp()).Methods("POST")
	r.Handle("/login", authHandler.Login()).Methods("POST")
}
