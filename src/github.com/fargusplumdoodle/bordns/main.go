package main

import (
	"fmt"
	"github.com/fargusplumdoodle/bordns/conf"
	"github.com/fargusplumdoodle/bordns/controller"
	"github.com/fargusplumdoodle/bordns/model"
	"github.com/gorilla/handlers"
	"log"
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
	err := http.ListenAndServe(conf.Env.ListenAddr, loggedRouter)
	if err != nil {
		log.Fatalf("server failed to start: %q", err)
	}
}
