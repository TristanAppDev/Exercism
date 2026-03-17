package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	lower := strings.ToLower(word)
	for index, char := range lower {
		if char != ' ' && char != '-' && index < strings.LastIndex(lower, string(char)) {
			return false
		}
	}
	return true
}
