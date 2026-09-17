// Package isbnverifier provides functions to the verification process is used to validate book identification numbers.
// These normally contain dashes and look like: 3-598-21508-8
package isbnverifier

// IsValidISBN validator ISBN
//
// The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only).
// In the case the check character is an X, this represents the value '10'.
// These may be communicated with or without hyphens, and can be checked for their validity
// by the following formula:
// (d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0
// If the result is 0, then it is a valid ISBN-10, otherwise it is invalid.
func IsValidISBN(isbn string) bool {
	pos := 0
	sum := 0
	for i := 0; i < len(isbn); i++ {
		c := isbn[i]
		if c == '-' {
			continue
		}
		if pos >= 10 {
			return false
		}
		var digit int
		switch {
		case c >= '0' && c <= '9':
			digit = int(c - '0')
		case c == 'X' && pos == 9:
			digit = 10
		default:
			return false
		}
		sum += digit * (10 - pos)
		pos++
	}
	return pos == 10 && sum%11 == 0
}
