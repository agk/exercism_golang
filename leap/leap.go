// Package leap provides functionality to determine whether a given year
// is a leap year according to the Gregorian calendar rules.
//
// The Gregorian calendar defines a leap year as:
//   - Every year that is evenly divisible by 4
//   - Except years that are evenly divisible by 100
//   - Unless the year is also evenly divisible by 400
//
// Examples:
//   - 1997 is not a leap year (not divisible by 4)
//   - 1900 is not a leap year (divisible by 100 but not by 400)
//   - 2000 is a leap year (divisible by 400)
package leap

// IsLeapYear reports whether the given year is a leap year in the
// Gregorian calendar.
//
// It returns true if the year is divisible by 4, except when it is
// divisible by 100 but not by 400. In other words:
//
//   year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
//
// Parameters:
//   - year: an integer representing the year (e.g., 2024, 1900)
//
// Returns:
//   - true if the year is a leap year, false otherwise
//
// Examples:
//   IsLeapYear(2000) // returns true
//   IsLeapYear(1900) // returns false
//   IsLeapYear(2024) // returns true
//   IsLeapYear(2023) // returns false
func IsLeapYear(year int) bool {
	return (year % 4 == 0) && (year % 100 != 0 || year % 400 == 0)
}