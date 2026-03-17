package proverb

import "fmt"

// The function Proverb takes a string array and creates a proverb out of it
func Proverb(rhyme []string) []string {
	lenR := len(rhyme)

	if lenR == 0 {
		return []string{}
	}

	res := make([]string, 0, lenR)
	for i := 0; i < lenR-1; i++ {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return res
}
