package controller

import "github.com/gorilla/mux"

var (
	domainController domain
)

func Startup(r *mux.Router) {
	domainController.registerRoutes(r)
}
