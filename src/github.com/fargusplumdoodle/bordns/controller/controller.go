package controller

import (
	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
)

var (
	domainController domain
	fqdnController   fqdn
)

func Startup(cli *clientv3.Client) *mux.Router {
	// Creating router
	r := mux.NewRouter()

	domainController.registerRoutes(r)
	fqdnController.registerRoutes(r, cli)

	return r
}
