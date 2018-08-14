package iot

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/edwardbrowncross/coffeemaker-server/data"
)

// CoffeeState handles interactions with the coffee machine IOT Thing Shadow document.
type CoffeeState struct {
	ThingName string
}

// NewCoffeeState creates a new coffee state object.
func NewCoffeeState(thingName string) CoffeeState {
	return CoffeeState{
		ThingName: thingName,
	}
}

// Get gets the coffee state from the reported field of the Thing Shaddow document.
func (cs *CoffeeState) Get() (state data.CoffeeState, err error) {
	sess, err := session.NewSession()
	if err != nil {
		return state, fmt.Errorf("failed to create AWS session: %v", err)
	}
	svc := iotdataplane.New(sess)
	req, res := svc.GetThingShadowRequest(&iotdataplane.GetThingShadowInput{ThingName: &cs.ThingName})

	err = req.Send()
	if err != nil {
		return state, fmt.Errorf("failed to get thing shadow: %v", err)
	}

	var shadow data.IOTState
	err = json.Unmarshal(res.Payload, &shadow)
	if err != nil {
		return state, fmt.Errorf("failed to unmarshall shadow document: %v", err)
	}
	state = shadow.Reported
	return
}