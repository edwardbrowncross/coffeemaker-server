package statechange

import (
	"github.com/edwardbrowncross/coffeemaker-server/data"
	log "github.com/sirupsen/logrus"
)

// Handler holds data required to handle a coffee state change.
type Handler struct {
	SlackSender func(string) error
}

// NewHandler creates a new Handler.
func NewHandler(ss func(string) error) Handler {
	return Handler{
		SlackSender: ss,
	}
}

// Handle handles a message sent to the shadow/update/acceted mqtt topic.
func (h *Handler) Handle(event data.Payload) error {
	state := event.State.Reported.CoffeeState
	log.WithField("state", state).Info("handling coffee state change")
	if state == "" {
		return nil
	}

	switch state {
	case "brewed":
		log.Info("sending brewed slack message")
		err := h.SlackSender("Coffee is ready!")
		if err != nil {
			log.WithError(err).Error("failed to send brewed slack message")
		}
	}
	log.Info("done")
	return nil
}
