package controller

import (
	"github.com/gorilla/mux"
)

var (
	domainController domain
	fqdnController   fqdn
)

func Startup() *mux.Router {
	// Creating router
	r := mux.NewRouter()

	domainController.registerRoutes(r)
	fqdnController.registerRoutes(r)

	return r
}
