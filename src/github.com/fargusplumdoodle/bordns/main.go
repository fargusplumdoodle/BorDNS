package main

import (
	"fmt"
	conf2 "github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/controller"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {
	// reading config
	conf := conf2.GetConfig()

	// connecting to etcd
	cli := conf2.SetupDB(conf.EtcdHosts)

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
