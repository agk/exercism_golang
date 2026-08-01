// Package isbnverifier provides functions for verifying ISBN-10 book identification numbers.
// ISBN-10 may contain hyphens and a check digit that can be a digit or 'X'.
package isbnverifier

import "strings"

// IsValidISBN validates an ISBN-10 string.
//
// The function accepts ISBN-10 with or without hyphens. The check digit can be
// a digit (0-9) or 'X' (representing 10) only in the last position.
//
// Validation formula:
//
//	(d1*10 + d2*9 + d3*8 + d4*7 + d5*6 + d6*5 + d7*4 + d8*3 + d9*2 + d10*1) % 11 == 0
//
// Examples:
//
//	IsValidISBN("3-598-21508-8")  // returns true
//	IsValidISBN("3-598-21507-X")  // returns true
//	IsValidISBN("3-598-21508-9")  // returns false
func IsValidISBN(isbn string) bool {
	// Remove hyphens as they are optional separators
	clean := strings.ReplaceAll(isbn, "-", "")

	// ISBN-10 must have exactly 10 characters after removing hyphens
	if len(clean) != 10 {
		return false
	}

	sum := 0
	for i, v := range clean {
		weight := 10 - i // weights: 10, 9, 8, ..., 1

		// 'X' is valid only as the check digit (last position) and represents 10
		if i == 9 && v == 'X' {
			sum += 10 * weight
			continue
		}

		// All other characters must be digits
		if v < '0' || v > '9' {
			return false
		}

		// Convert digit character to its numeric value
		sum += int(v-'0') * weight
	}

	// Valid ISBN-10 sum is divisible by 11
	return sum%11 == 0
}
