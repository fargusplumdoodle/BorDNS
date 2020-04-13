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
	client := model.SetupDB(conf.Env.EtcdHosts)
	defer client.Close()

	// start controllers
	r := controller.Startup()

	// logging requests
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// running server!
	fmt.Println("Starting BorDNS API on port 8000, connected to etcd", client.Endpoints())
	http.ListenAndServe(conf.Env.ListenAddr, loggedRouter)

	// this should never execute, im still learning this all
	fmt.Println("Exiting! There was probably an error")
}
