package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type domain struct{}

func (d domain) registerRoutes(r *mux.Router) {
	r.HandleFunc("/domain", handleDomains)
}

func handleDomains(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("eyyyy"))
}
