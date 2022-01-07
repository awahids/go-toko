package main

import (
	"log"
	"net/http"

	"github.com/awahids/tokoDistributor/database"
	"github.com/awahids/tokoDistributor/handlers"
	"github.com/awahids/tokoDistributor/models"
	"github.com/gorilla/mux"
)

func handlersRequest()  {
	log.Println("server running at http:localhost:8000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Welcome)
	router.HandleFunc("/api/v1/signup", handlers.AddUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//database setUp
	db := database.Database()
	db.AutoMigrate(&models.User{})
	handlersRequest()
}