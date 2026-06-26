package raindrops

import (
	"strconv"
	"strings"
)

// table rules: divider -> sound
var rules = []struct {
	divider int
	sound   string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	sb := strings.Builder{}

	for _, rule := range rules {
		if number%rule.divider == 0 {
			sb.WriteString(rule.sound)
		}
	}

	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(number)) // If the rules don't work
	}
	return sb.String()
}
