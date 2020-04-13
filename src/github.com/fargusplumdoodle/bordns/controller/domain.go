package controller

import (
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/model"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"net/http"
)

type domain struct{}

func (d domain) registerRoutes(r *mux.Router) {
	r.Handle("/domain", httpauth.SimpleBasicAuth(conf.Env.AuthUsername, conf.Env.AuthPassword)(d))
}

func (d domain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get all domains
	allDomains, err := model.GetAllDomains()

	if err != nil {
		respondBadRequest(w, err)
	}

	// encoding into json
	encodeAllDomains(w, allDomains)

}
