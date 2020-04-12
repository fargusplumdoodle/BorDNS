package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
	"net"
	"net/http"
	"regexp"
)

type fqdn struct{}

// etcd client
var cli *clientv3.Client

const (
	DOMAIN_REGEX = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9])).([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}.[a-zA-Z]{2,3})$`
	IPV4_REGEX   = `^(([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.){3}([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])$`
)

var domainPattern, _ = regexp.Compile(DOMAIN_REGEX)
var ipPattern, _ = regexp.Compile(IPV4_REGEX)

func (f fqdn) registerRoutes(r *mux.Router, etcdClient *clientv3.Client) {
	cli = etcdClient

	r.HandleFunc("/fqdn", handleFQDNS)
}
func handleFQDNS(w http.ResponseWriter, r *http.Request) {
	// TODO: validation
	// Validation
	err := validateRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(
			"bad request: %q", err.Error(),
		)))
	}

	switch r.Method {
	case http.MethodGet:
		// TODO: retrieve IP of fqdn
	case http.MethodPost:
		// TODO: validate, must contain IP field
		// TODO: create A record
		ctx, cancel := context.WithTimeout(context.Background(), conf.DB_TIMEOUT)
		resp, err := cli.Put(ctx, "/bor/bor/test", `{"host":"10.0.0.1","ttl":60}`)
		cancel()
		if err != nil {
			fmt.Errorf("Error adding dns name to etcd, %q", err)
		}
		fmt.Printf("response:, %q", resp)
	case http.MethodDelete:
		// TODO: delete the A record
	default:
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
func validateRequest(r *http.Request) error {
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
		return errors.New("invalid method")
	}

	// All fields have fqdn
	FqdnParams := queryString["FQDN"]
	if len(FqdnParams) == 0 {
		return errors.New("request did not contain FQDN parameter")
	}
	// checking fqdn is valid
	matches := domainPattern.FindStringSubmatch(FqdnParams[0])
	if len(matches) == 0 {
		return errors.New("invalid FQDN parameter")
	}

	// the post method requires the an IP
	if r.Method == http.MethodPost {
		ipParams := queryString["IP"]
		if len(ipParams) == 0 {
			return errors.New("request did not contain IP parameter")
		}
		// ensuring IP is valid
		if net.ParseIP(ipParams[0]) == nil {
			return errors.New("invalid IP parameter")
		}
	}

	return nil
}
