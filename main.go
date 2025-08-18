package main

import (
	"fmt"
	"net/http"

	"github.com/gmayas/ApiRestGo/src/db"
	"github.com/gmayas/ApiRestGo/src/models"
	"github.com/gmayas/ApiRestGo/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	//Conection DB
	db.DBConnection()
	// Execute Models
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	// Create a new router
	r := mux.NewRouter()
	// Define the home route
	r.HandleFunc("/", routes.HomeHandler).Methods("GET")
	// Start the server, port 3000
	fmt.Println("ApiRestGo server run port: 3000.")
	http.ListenAndServe(":3000", r)

}
