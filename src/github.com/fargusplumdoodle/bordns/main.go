package main

import (
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/controller"
	"github.com/fargusplumdoodle/bordns/model"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {
	// reading config
	conf.SetupConfig()

	// connecting to etcd
	model.SetupDB(conf.Env.EtcdHosts)

	// start controllers
	r := controller.Startup()

	// logging requests
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// running server!
	fmt.Println("Starting BorDNS API on port 8000")
	http.ListenAndServe(conf.Env.ListenAddr, loggedRouter)

	// this should never execute, im still learning this all
	fmt.Println("Exiting! There was probably an error")
}
