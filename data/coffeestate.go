package data

import "time"

const coffeePortion = 200

// CoffeeState holds the parameters sent by the coffee machine to report its state.
type CoffeeState struct {
	Heating      bool      `json:"heaterOn"`
	JugPresent   bool      `json:"jugPresent"`
	LightState   string    `json:"light"`
	CoffeeState  string    `json:"coffee"`
	CoffeeWeight int64     `json:"coffeeWeight"`
	LastBrewTime time.Time `json:"lastBrewTime"`
}

// GetCupsRemaining returns the number of cups left in the jug
func (c *CoffeeState) GetCupsRemaining() int {
	return int((c.CoffeeWeight + coffeePortion/2) / coffeePortion)
}

// GetTimeSinceLastBrew returns the duration since the coffee machine was last tuned on.
func (c *CoffeeState) GetTimeSinceLastBrew() time.Duration {
	return time.Now().Sub(c.LastBrewTime)
}
