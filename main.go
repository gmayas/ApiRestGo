package main

import (
	"fmt"
	"net/http"
	"github.com/gmayas/ApiRestGo/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()
	// Define the home route
	r.HandleFunc("/", routes.HomeHandler).Methods("GET")
	// Start the server, port 3000
	fmt.Println("ApiRestGo server run port: 3000.")
	http.ListenAndServe(":3000", r)

}
