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
	// Define user routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	// Define task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	// Start the server, port 3000
	fmt.Println("ApiRestGo server run port: 3000.")
	http.ListenAndServe(":3000", r)
}
