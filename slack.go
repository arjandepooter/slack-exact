package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type slackRequest struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	IconURL  string `json:"icon_url,omit_empty"`
	Channel  string `json:"channel"`
}

func postToSlack(command *SlackCommand, message string) error {
	payload := &slackRequest{
		Channel:  fmt.Sprintf("#%s", command.ChannelName),
		Username: config.BotUsername,
		Text:     fmt.Sprintf("%s %s\n%s", command.Command, command.Text, message),
		IconURL:  fmt.Sprintf("%s/exact_online.png", config.BaseURL),
	}

	request := gorequest.New()
	resp, body, errs := request.Post(config.WebhookURL).Send(*payload).End()
	if len(errs) > 0 {
		return fmt.Errorf("Invalid request: %s", errs[0].Error())
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Invalid request: %s", body)
	}

	return nil
}
