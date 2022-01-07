package main

import (
  "fmt"
	"log"
	"net/http"

	// "github.com/awahids/tokoDistributor/database" //run localhost
	// "github.com/awahids/tokoDistributor/entity" //run localhost
  "github.com/replit/database-go" //run at replit.com
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
	// db := database.Database()
	// db.AutoMigrate(&entity.Users{})
  
  database.Set("key", "value")
  val, _ := database.Get("key")
  fmt.Println(val)
	handlersRequest()
}