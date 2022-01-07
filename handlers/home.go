package handlers

import "net/http"

func Welcome(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("welcome!\n"))
}
