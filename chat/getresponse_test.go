package chat

import (
	"testing"
	"time"

	"github.com/edwardbrowncross/coffeemaker-server/data"
)

func TestGetResponse(t *testing.T) {
	tests := []struct {
		input    data.CoffeeState
		response string
	}{
		{
			input: data.CoffeeState{
				CoffeeState:  "off",
				JugPresent:   true,
				CoffeeWeight: 400,
			},
			response: "Coffee machine is off. Why not make some coffee?",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "stale",
				JugPresent:   true,
				CoffeeWeight: 90,
			},
			response: "There is no coffee left. Why not make some more?",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "stale",
				JugPresent:   true,
				CoffeeWeight: 490,
				LastBrewTime: time.Now().Add(-2 * time.Hour),
			},
			response: "There are 2 cups of stale coffee left. It was made 2 hours ago.",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "brewed",
				JugPresent:   true,
				Heating:      true,
				CoffeeWeight: 200,
				LastBrewTime: time.Now().Add(-10 * time.Minute),
			},
			response: "There is 1 cup of fresh coffee left. It was made just now.",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "stale",
				JugPresent:   true,
				Heating:      true,
				CoffeeWeight: 300,
				LastBrewTime: time.Now().Add(-24 * time.Hour),
			},
			response: "There is only stale coffee from yesterday. Why not make some fresh?",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "brewed",
				JugPresent:   false,
				Heating:      true,
				CoffeeWeight: 800,
				LastBrewTime: time.Now().Add(-30 * time.Minute),
			},
			response: "There were 4 cups of fresh coffee left, but someone's taking some right now. Don't delay!",
		},
		{
			input: data.CoffeeState{
				CoffeeState: "preparing",
				JugPresent:  false,
			},
			response: "Someone is making coffee right now.",
		},
		{
			input: data.CoffeeState{
				CoffeeState:  "brewing",
				LastBrewTime: time.Now().Add(-5 * time.Minute),
			},
			response: "Coffee is brewing. It will be ready in 4 minutes.",
		},
	}

	for _, test := range tests {
		res := GetResponse(test.input)
		if res != test.response {
			t.Errorf("Expected response of %s but got %s", test.response, res)
		}
	}
}
