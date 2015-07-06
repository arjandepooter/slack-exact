package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/slack", CommandHandler).Methods("POST")

	http.ListenAndServe(":3000", router)
}
