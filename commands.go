package main

import (
	"fmt"
	"github.com/mattn/go-shellwords"
)

// SlackCommand represents the payload of the Slack slash commands
type SlackCommand struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Command     string `schema:"command"`
	Text        string `schema:"text"`
}

func (command *SlackCommand) getParsedText() []string {
	parsedText, err := shellwords.NewParser().Parse(command.Text)
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

func helpCommand(*SlackCommand) string {
	return "Use one of the following commands:\n\tversion\n\thelp"
}

func versionCommand(*SlackCommand) string {
	return fmt.Sprintf("SlackExact version %s", VERSION)
}
