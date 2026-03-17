package luhn

import (
	"strings"
)

// Valid checks if a string is a valid credit card number
func Valid(id string) bool {
	if len(strings.TrimSpace(id)) <= 1 {
		return false
	}

	counter := 0
	sum := 0

	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}

		digit := int(r - '0')
		if counter%2 == 1 {
			sum += doubled(digit)
		} else {
			sum += digit
		}
		counter++
	}
	return counter > 1 && sum%10 == 0
}

func doubled(x int) int {
	x *= 2
	if x > 9 {
		x -= 9
	}
	return x
}
