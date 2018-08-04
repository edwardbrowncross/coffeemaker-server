package data

// Payload holds the contents of a AWS IOT uodate object.
type Payload struct {
	State     IOTState `json:"state"`
	Version   int64    `json:"version"`
	Timestamp int64    `json:"timestamp"`
}

// IOTState is the state field of an AWS IOT update object.
type IOTState struct {
	Reported CoffeeState `json:"reported"`
	Desired  CoffeeState `json:"desired"`
}
