package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type configSpecification struct {
	// Port
	Port int `default:"3000"`
	// URL of the Slack Webhook
	WebhookURL string `envconfig:"webhook_url"`
	// Token for incoming Slack requests
	IncomingToken string `envconfig:"incoming_token"`
	// Username for bot replies
	BotUsername string `envconfig:"bot_username" default:"Exact Online"`
	// Base URL of Slack Exact
	BaseURL string `envconfig:"base_url" default:"http://slackexact.greymug.io"`
	// Exact app key
	AppKey string `envconfig:"app_key"`
	// Exact client ID
	ClientID string `envconfig:"client_id"`
	// Exact client secret
	ClientSecret string `envconfig:"client_secret"`
}

var config configSpecification

func init() {
	err := envconfig.Process("slack_exact", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
