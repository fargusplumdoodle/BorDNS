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
	r := mux.NewRouter()
	// TODO: read environment variables
	// TODO: setup connection to etcd
	// start controllers
	controller.Startup(r)

	// logging requests
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	//
	fmt.Println("Starting BorDNS API on port 8000")
	http.ListenAndServe(":8000", loggedRouter)

	// this should never execute
	fmt.Println("Exiting! There was probably an error")
}
