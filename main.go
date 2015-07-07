package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Current version
const VERSION = "0.0.1"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/slack", commandHandler).Methods("POST")

	http.ListenAndServe(":3000", router)
}
