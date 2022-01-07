package handlers

import "net/http"

func YourHandler(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
