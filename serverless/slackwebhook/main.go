package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edwardbrowncross/coffeemaker-server/chat"
	"github.com/edwardbrowncross/coffeemaker-server/config"
	"github.com/edwardbrowncross/coffeemaker-server/iot"
	"github.com/edwardbrowncross/coffeemaker-server/slackwebhook"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	c := config.FromEnvironment()
	cs := iot.NewCoffeeState(c.ThingName)
	h := slackwebhook.NewHandler(cs.Get, chat.GetResponse)
	lambda.Start(h.Handle)
}
