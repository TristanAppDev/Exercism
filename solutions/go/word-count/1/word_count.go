package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	phrase = strings.Replace(phrase, ",", " ", -1)
	phrase = strings.Replace(phrase, ".", " ", -1)
	phrase = strings.Replace(phrase, ":", " ", -1)
	words := strings.Fields(phrase)

	for i := 0; i < len(words); i++ {
		words[i] = strings.Trim(words[i], ".,!$%^&*()_+-=<>?/|:;\"'@")
	}

	freq := make(Frequency)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
