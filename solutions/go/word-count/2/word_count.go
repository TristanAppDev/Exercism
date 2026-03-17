package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var reg = regexp.MustCompile(`[^a-zA-Z0-9']+`)

func WordCount(phrase string) Frequency {
	words := strings.Fields(normalize(phrase))
	freq := Frequency{}
	for _, word := range words {
		word = strings.Trim(word, "'")
		freq[word]++
	}
	return freq
}

func normalize(phrase string) string {
	phrase = strings.ToLower(phrase)
	return reg.ReplaceAllString(phrase, " ")
}
