// Package weather: This package gives the weather forecast for the specified city.
package weather

// CurrentCondition: is the current weather condition in the city.
var CurrentCondition string

// CurrentLocation: is the city where the weather forecast is being made.
var CurrentLocation string

// Forecast: Returns the current weather condition for the specified city as a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
