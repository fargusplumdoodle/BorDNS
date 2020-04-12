package main

import (
	"fmt"
	"github.com/fargusplumdoodle/bordns/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	// reading config
	conf := getConfig()

	// TODO: setup connection to etcd
	cli := setupDB(conf.EtcdHosts)

	// start controllers
	r := controller.Startup(cli)

	// logging requests
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// running server!
	fmt.Println("Starting BorDNS API on port 8000")
	http.ListenAndServe(conf.ListenAddr, loggedRouter)

	// this should never execute, im still learning this all
	fmt.Println("Exiting! There was probably an error")
}
