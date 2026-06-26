// Package weather comment.
package weather

var (
    // CurrentCondition variables comment.
    CurrentCondition string
    // CurrentLocation variables comment.
    CurrentLocation  string
)
// Forecast() function does ----.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
