package pangram

import "strings"

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// IsPangram returns true if the input string contains all the letters of the alphabet at least once.
func IsPangram(input string) bool {
	if input == "" {
		return false
	}

	input = strings.ToLower(input)

	return All(alphabet, func(letter string) bool {
		return strings.Contains(input, letter)
	})
}

// All returns true if all elements in the list return true for the predicate.
func All(elements []string, predicate func(string) bool) bool {
	for _, element := range elements {
		if !predicate(element) {
			return false
		}
	}
	return true
}
