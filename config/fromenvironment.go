package config

import "os"

// Config holds config data for the application.
type Config struct {
	ThingName    string
	SlackWebhook string
}

func FromEnvironment() Config {
	return Config{
		ThingName:    os.Getenv("THING_NAME"),
		SlackWebhook: os.Getenv("SLACK_WEBHOOK"),
	}
}
