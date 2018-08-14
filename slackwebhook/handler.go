package slackwebhook

import (
	"context"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/edwardbrowncross/coffeemaker-server/data"
	log "github.com/sirupsen/logrus"
)

// Handler handles a webhook call.
type Handler struct {
	StateGetter     func() (data.CoffeeState, error)
	ResponseBuilder func(data.CoffeeState) string
}

// NewHandler creates a new handler.
func NewHandler(sg func() (data.CoffeeState, error), rb func(data.CoffeeState) string) Handler {
	return Handler{
		StateGetter:     sg,
		ResponseBuilder: rb,
	}
}

// Handle is invoked by the lambda on a web request.
func (h *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	args, err := url.ParseQuery(req.Body)
	if err != nil {
		log.WithError(err).Error("failed to decode request body")
		return
	}
	log.WithField("args", args).Info("handling slack webhook")

	state, err := h.StateGetter()
	if err != nil {
		log.WithError(err).Error("failed to fetch thing shadow")
		return
	}

	s := h.ResponseBuilder(state)

	res.StatusCode = http.StatusOK
	res.Body = s
	log.WithField("state", state).WithField("response", s).Info("done")
	return
}
