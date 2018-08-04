package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edwardbrowncross/coffeemaker-server/config"
	"github.com/edwardbrowncross/coffeemaker-server/slack"
	"github.com/edwardbrowncross/coffeemaker-server/statechange"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	c := config.FromEnvironment()
	ss := slack.NewSender(c.SlackWebhook)
	h := statechange.NewHandler(ss.Send)
	lambda.Start(h.Handle)
}
