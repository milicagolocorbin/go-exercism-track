// Package weather provides tools for weather information.
package weather

// CurrentCondition represents a certain condition of the weather.
var CurrentCondition string

// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns a string value that shows conditions of the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
