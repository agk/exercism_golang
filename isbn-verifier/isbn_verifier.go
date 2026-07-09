// Package isbnverifier provides functions to the verification process is used to validate book identification numbers.
// These normally contain dashes and look like: 3-598-21508-8
package isbnverifier

import (
	"strconv"
	"strings"
)

// IsValidISBN validator ISBN
//
// The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only).
// In the case the check character is an X, this represents the value '10'.
// These may be communicated with or without hyphens, and can be checked for their validity
// by the following formula:
// (d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0
// If the result is 0, then it is a valid ISBN-10, otherwise it is invalid.
func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
	if len(s) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		var digit int
		var err error

		// The last character can be 'X'
		if i == 9 && s[i] == 'X' {
			digit = 10
		} else {
			digit, err = strconv.Atoi(string(s[i]))
			if err != nil {
				return false
			}
		}

		// Multiplied by weight (10, 9, 8, ..., 1)
		sum += digit * (10 - i)
	}

	return sum%11 == 0
}
