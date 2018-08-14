package config

import "os"

// Config holds config data for the application.
type Config struct {
	Region       string
	IOTEndpoint  string
	ThingName    string
	SlackWebhook string
}

func FromEnvironment() Config {
	return Config{
		Region:       os.Getenv("AWS_REGION"),
		IOTEndpoint:  os.Getenv("IOT_ENDPOINT"),
		ThingName:    os.Getenv("THING_NAME"),
		SlackWebhook: os.Getenv("SLACK_WEBHOOK"),
	}
}
