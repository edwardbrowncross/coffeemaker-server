package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Sender holds data required to send a slack notification.
type Sender struct {
	WebhookURL string
}

// NewSender creates a new Sender.
func NewSender(webhook string) Sender {
	return Sender{
		WebhookURL: webhook,
	}
}

// Send sends a slack message to the webhook with the given message.
func (s *Sender) Send(m string) (err error) {
	body := map[string]string{
		"text": m,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = http.Post(s.WebhookURL, "application/json", bytes.NewBuffer(jsonBody))
	return err
}
