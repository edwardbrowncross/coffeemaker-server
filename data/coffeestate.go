package data

import "time"

// CoffeeState holds the parameters sent by the coffee machine to report its state.
type CoffeeState struct {
	Heating      bool      `json:"heating"`
	CoffeeState  string    `json:"coffee"`
	LastBrewTime time.Time `json:"lastBrewTime"`
}
