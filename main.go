package main

import (
	"log"
	"net/http"

	"github.com/awahids/tokoDistributor/database"
	"github.com/awahids/tokoDistributor/entity"
	"github.com/awahids/tokoDistributor/handlers"
	"github.com/gorilla/mux"
)

func handlersRequest()  {
	log.Println("server running at http:localhost:8000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.YourHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//database setUp
	db := database.Database()
	db.AutoMigrate(&entity.Users{})
	handlersRequest()
}