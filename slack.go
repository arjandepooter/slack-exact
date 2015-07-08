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

func newSlackRequest() *slackRequest {
	req := new(slackRequest)
	req.IconURL = fmt.Sprintf("%s/exact_online.png", config.BaseURL)
	req.Username = config.BotUsername

	return req
}

func postToSlack(payload *slackRequest) error {
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
