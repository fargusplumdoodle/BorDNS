package controller

import (
	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
)

var (
	domainController domain
)

func Startup(cli *clientv3.Client) *mux.Router {
	r := mux.NewRouter()

	domainController.registerRoutes(r)
	println(cli)

	return r
}
