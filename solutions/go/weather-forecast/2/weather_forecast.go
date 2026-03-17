// Package weather: This package gives the weather forecast for the specified city
package weather

// CurrentCondition: Is the current weather
var CurrentCondition string

// CurrentLocation: Location of the current weather forecast
var CurrentLocation string

// Forecast: Returns the current weather forecast for a specified city as a string
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
