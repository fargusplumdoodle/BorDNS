package controller

import (
	"errors"
	"fmt"
	"github.com/fargusplumdoodle/bordns/model"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

type fqdn struct{}

const (
	DOMAIN_REGEX = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9])).([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}.[a-zA-Z]{2,3})$`
	IPV4_REGEX   = `^(([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.){3}([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])$`
)

var domainPattern, _ = regexp.Compile(DOMAIN_REGEX)

func (f fqdn) registerRoutes(r *mux.Router) {
	r.HandleFunc("/fqdn", handleFQDNS)
}

func handleFQDNS(w http.ResponseWriter, r *http.Request) {
	// Validation
	queryString, err := validateRequest(r)
	if err != nil {
		respondBadRequest(w, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// TODO: retrieve IP of fqdn
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
		}
	case http.MethodDelete:
		// TODO: delete the A record
	default:
	}

}

func respondBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(
		"bad request: %q", err.Error(),
	)))
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
