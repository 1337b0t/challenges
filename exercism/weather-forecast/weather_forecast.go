// Package weather provides the current city weather conditions.
package weather

// CurrentCondition stores the weather condition as a string.
var CurrentCondition string

// CurrentLocation stores the current city location as a string.
var CurrentLocation string

// Forecast returns a string format of the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
