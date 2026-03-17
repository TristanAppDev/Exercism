package luhn

import (
	"strings"
)

func Valid(id string) bool {
	if len(strings.TrimSpace(id)) <= 1 {
		return false
	}

	digits := make([]int, 0, len(id))

	for _, r := range id {
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		digits = append(digits, int(r-'0'))
	}

	return sum(digits)%10 == 0
}

func sum(digits []int) int {
	sum := 0
	for i, digit := range digits {
		if dd, ok := doubled(digit, i, len(digits)); ok {
			digit = dd
		}
		sum += digit
	}
	return sum
}

func doubled(x, i, length int) (int, bool) {
	x *= 2
	if x > 9 {
		x -= 9
	}
	return x, (length-i)%2 == 0
}
