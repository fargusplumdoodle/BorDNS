package controller

import (
	"encoding/json"
	"fmt"
	"github.com/fargusplumdoodle/bordns/viewmodel"
	"github.com/gorilla/mux"
	"net/http"
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

func respondBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(
		"bad request: %q", err.Error(),
	)))
}

func encodeARecord(w http.ResponseWriter, fqdn, ip string) {
	/*
		Writes A record response
	*/
	aRecord := viewmodel.Arecord{
		IP:   ip,
		FQDN: fqdn,
	}

	// encoding A record
	enc := json.NewEncoder(w)
	err := enc.Encode(&aRecord)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
}

func encodeAllDomains(w http.ResponseWriter, mp []viewmodel.Zone) {
	/*
		Writes a list of zones to  A records
	*/
	data, err := json.Marshal(mp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
