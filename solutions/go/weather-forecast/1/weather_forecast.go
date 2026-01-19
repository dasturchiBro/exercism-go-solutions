// Package weather to get information about the weather forecasts.
package weather

var (
// CurrentCondition receives the current weather condition in the location provided by CurrentLocation.
	CurrentCondition string
// CurrentLocation receives the current location.
	CurrentLocation  string
)

// Forecast function return the current weather condition in a particular location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
