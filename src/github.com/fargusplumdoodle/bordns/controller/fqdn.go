package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type fqdn struct{}

func (f fqdn) registerRoutes(r *mux.Router) {
	r.HandleFunc("/fqdn", handleFQDNS)
}
func handleFQDNS(w http.ResponseWriter, r *http.Request) {
	// TODO: validation
	//  All requests must contain FQDN query parameter
	//  with a valid FQDN

	switch r.Method {
	case http.MethodGet:
		// TODO: retrieve IP of fqdn
	case http.MethodPost:
		// TODO: validate, must contain IP field
		// TODO: create A record
	case http.MethodDelete:
		// TODO: delete the A record
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
	}

}
