package chat

import (
	"fmt"
	"time"

	"github.com/edwardbrowncross/coffeemaker-server/data"
)

// GetResponse returns the human-friendly description of the coffee state.
func GetResponse(coffee data.CoffeeState) string {
	if coffee.CoffeeState == "preparing" {
		return "Someone is making coffee right now."
	} else if coffee.CoffeeState == "brewing" {
		return fmt.Sprintf("Coffee is brewing. It will be ready in %d minutes.", 9-coffee.GetTimeSinceLastBrew()/time.Minute)
	} else if coffee.GetCupsRemaining() < 1 {
		return "There is no coffee left. Why not make some more?"
	} else if coffee.GetTimeSinceLastBrew()/time.Hour > 12 {
		cd := getCoffeeDescription(coffee.CoffeeState)
		ts := getTimeString(coffee.LastBrewTime)
		return fmt.Sprintf("There is only %s coffee from %s. Why not make some fresh?", cd, ts)
	} else if !coffee.JugPresent {
		ww := getWereWas(coffee.GetCupsRemaining())
		cs := getCups(coffee.GetCupsRemaining())
		cd := getCoffeeDescription(coffee.CoffeeState)
		return fmt.Sprintf("There %s %s of %s coffee left, but someone's taking some right now. Don't delay!", ww, cs, cd)
	}
	ia := getIsAre(coffee.GetCupsRemaining())
	cs := getCups(coffee.GetCupsRemaining())
	cd := getCoffeeDescription(coffee.CoffeeState)
	ts := getTimeString(coffee.LastBrewTime)
	return fmt.Sprintf("There %s %s of %s coffee left. It was made %s.", ia, cs, cd, ts)
}

func getIsAre(n int) string {
	if n == 1 {
		return "is"
	}
	return "are"
}

func getWereWas(n int) string {
	if n == 1 {
		return "was"
	}
	return "were"
}

func getCups(n int) string {
	if n == 1 {
		return "1 cup"
	} else {
		return fmt.Sprintf("%d cups", n)
	}
}

func getCoffeeDescription(s string) string {
	if s == "brewed" {
		return "fresh"
	} else if s == "stale" {
		return "stale"
	} else if s == "reheating" {
		return "reheated"
	}
	return ""
}

func getTimeString(t time.Time) string {
	dt := time.Now().Sub(t)
	var hours = dt / time.Hour
	if dt/time.Minute < 15 {
		return "just now"
	} else if hours < 1 {
		return "within the last hour"
	} else if hours < 12 {
		return fmt.Sprintf("%d hours ago", hours)
	} else if hours < 48 {
		return "yesterday"
	} else if hours >= 168 || time.Now().Weekday() < t.Weekday() {
		return "last week"
	} else {
		return fmt.Sprintf("%d days ago", hours/24)
	}
}
