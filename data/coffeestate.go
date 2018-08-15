package data

import "time"

// CoffeeState holds the parameters sent by the coffee machine to report its state.
type CoffeeState struct {
	Heating      bool      `json:"heaterOn"`
	JugPresent   bool      `json:"jugPresent"`
	LightState   string    `json:"light"`
	CoffeeState  string    `json:"coffee"`
	CoffeeWeight int64     `json:"coffeeWeight"`
	LastBrewTime time.Time `json:"lastBrewTime"`
}

func (c *CoffeeState) GetCupsRemaining() int {
	return int((c.CoffeeWeight + 125) / 250.0)
}

func (c *CoffeeState) GetTimeSinceLastBrew() time.Duration {
	return time.Now().Sub(c.LastBrewTime)
}
