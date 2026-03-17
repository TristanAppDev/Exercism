package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)

	unique := distinct(strings.Split(word, ""))

	return len(unique) == len(word)
}

func distinct(slice []string) string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return strings.Join(list, "")
}
