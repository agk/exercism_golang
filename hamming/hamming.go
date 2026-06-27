package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	cnt := 0
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return cnt, errors.New("Lengths differents")
	}

	for index := range a {
		if rune(a[index]) != rune(b[index]) {
			cnt++
		}
	}
	return cnt, nil
}
