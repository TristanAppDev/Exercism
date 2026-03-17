package scrabble

import "strings"

func Score(word string) int {
	score := 0
	word = strings.ToUpper(word)
	for _, l := range word {
		score += pointsPerLetter(l)
	}
	return score
}

func pointsPerLetter(l rune) int {
	switch l {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	default:
		return 10
	}
}
