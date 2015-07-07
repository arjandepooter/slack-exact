package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// Current version
const VERSION = "0.0.1"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/slack", commandHandler).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%d", config.Port))
}
