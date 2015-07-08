package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var userStore *UserStore

func main() {
	userStore = new(UserStore)

	router := mux.NewRouter()
	router.HandleFunc("/slack", commandHandler).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%d", config.Port))
}
