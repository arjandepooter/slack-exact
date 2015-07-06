package main

import (
	"github.com/gorilla/schema"
	"io"
	"net/http"
)

func CommandHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	command := new(SlackCommand)
	decoder := schema.NewDecoder()

	err := decoder.Decode(command, req.PostForm)
	if err != nil {
		res.WriteHeader(400)
		io.WriteString(res, "Invalid form data\n")
		return
	}
}
