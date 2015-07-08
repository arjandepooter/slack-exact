package main

import (
	"github.com/gorilla/schema"
	"io"
	"net/http"
)

var commands = map[string](func(*SlackCommand, *User) error){
	"help":    helpCommand,
	"version": versionCommand,
}

func commandHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	command := new(SlackCommand)
	decoder := schema.NewDecoder()

	err := decoder.Decode(command, req.PostForm)
	if err != nil {
		res.WriteHeader(400)
		io.WriteString(res, "Invalid form data\n")
		return
	}

	if command.Token != config.IncomingToken {
		res.WriteHeader(403)
		io.WriteString(res, "Invalid token\n")
		return
	}

	user, created := userStore.GetOrAdd(command.UserID)
	if created {
		user.UserName = command.UserName
	}

	commandFunc, exists := commands[command.getSubCommand()]
	if !exists {
		commandFunc = commands["help"]
	}
	err = commandFunc(command, user)
	if err != nil {
		res.WriteHeader(400)
		io.WriteString(res, err.Error())
	}
}
