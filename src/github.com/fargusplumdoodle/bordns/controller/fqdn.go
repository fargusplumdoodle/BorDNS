package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/model"
	"github.com/fargusplumdoodle/bordns/viewmodel"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

type fqdn struct{}

const (
	DOMAIN_REGEX = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9])).([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}.[a-zA-Z]{2,3})$`
)

var domainPattern, _ = regexp.Compile(DOMAIN_REGEX)

func (f fqdn) registerRoutes(r *mux.Router) {
	r.Handle("/fqdn", httpauth.SimpleBasicAuth(conf.Env.AuthUsername, conf.Env.AuthPassword)(f))
}

func (f fqdn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validation
	queryString, err := validateRequest(r)
	if err != nil {
		respondBadRequest(w, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		ip, err := model.GetCoreDNSRecordForHost(queryString["FQDN"][0])
		if err != nil {
			respondBadRequest(w, err)
			return
		} else {
			// This should probably be JSON encoded
			w.WriteHeader(http.StatusOK)
			encodeARecord(w, queryString["FQDN"][0], ip.Host)

		}
	case http.MethodPost:
		err = model.AddARecord(
			queryString["FQDN"][0],
			queryString["IP"][0],
		)
		if err != nil {
			respondBadRequest(w, err)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			encodeARecord(w, queryString["FQDN"][0], queryString["IP"][0])
		}
	case http.MethodDelete:
		err = model.DeleteARecord(
			queryString["FQDN"][0],
		)
		if err != nil {
			respondBadRequest(w, err)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}

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

/*
	Validates request
	-----------------

	Checks query string for the following
    attributes based on the request method

	POST:
      - FQDN
      - IP
     GET:
      - FQDN
     DELETE:
      - FQDN

	Any other method, is a bad request
*/
func validateRequest(r *http.Request) (url.Values, error) {
	// getting query string from request
	queryString := r.URL.Query()
	acceptedMethods := []string{http.MethodGet, http.MethodDelete, http.MethodPost}

	// checking valid method
	validMethod := false
	for _, method := range acceptedMethods {
		if r.Method == method {
			validMethod = true
			break
		}
	}
	if !validMethod {
		return nil, errors.New("invalid method")
	}

	// All fields have fqdn
	FqdnParams := queryString["FQDN"]
	if len(FqdnParams) == 0 {
		return nil, errors.New("request did not contain FQDN parameter")
	}
	// checking fqdn is valid
	matches := domainPattern.FindStringSubmatch(FqdnParams[0])
	if len(matches) == 0 {
		return nil, errors.New("invalid FQDN parameter")
	}

	// the post method requires the an IP
	if r.Method == http.MethodPost {
		ipParams := queryString["IP"]
		if len(ipParams) == 0 {
			return nil, errors.New("request did not contain IP parameter")
		}
		// ensuring IP is valid
		if net.ParseIP(ipParams[0]) == nil {
			return nil, errors.New("invalid IP parameter")
		}
	}

	return queryString, nil
}
