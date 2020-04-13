package controller

import (
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"net/http"
)

type domain struct{}

func (d domain) registerRoutes(r *mux.Router) {
	r.Handle("/domain", httpauth.SimpleBasicAuth(conf.Env.AuthUsername, conf.Env.AuthPassword)(d))
}

func (d domain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("TODO: show list of domains that are registered"))
	// TODO: show list of domains that are registered
}
