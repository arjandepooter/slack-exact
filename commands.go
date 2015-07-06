package main

import (
	"github.com/mattn/go-shellwords"
)

// SlackCommand represents the payload of the Slack slash commands
type SlackCommand struct {
	token       string `schema:"token"`
	teamID      string `schema:"team_id"`
	teamDomain  string `schema:"team_domain"`
	channelID   string `schema:"channel_id"`
	channelName string `schema:"channel_name"`
	userID      string `schema:"user_id"`
	userName    string `schema:"user_name"`
	command     string `schema:"command"`
	text        string `schema:"text"`
}

func (command *SlackCommand) getParsedText() []string {
	parsedText, err := shellwords.NewParser().Parse(command.text)
	if err != nil {
		return []string{}
	}
	return parsedText
}

func (command *SlackCommand) getSubCommand() string {
	parsedText := command.getParsedText()

	if len(parsedText) > 0 {
		return parsedText[0]
	}
	return "help"
}
