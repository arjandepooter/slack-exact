package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port int `default:"3000"`
}

var config Config

func init() {
	err := envconfig.Process("slack_exact", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
