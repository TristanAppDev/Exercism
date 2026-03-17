package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, v := range candidates {
		if isAnagram(subject, v) {
			anagrams = append(anagrams, v)
		}
	}
	return anagrams
}

func isAnagram(subject, candidate string) bool {

	if len(subject) != len(candidate) {
		return false
	}

	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)

	if subject == candidate {
		return false
	}

	for _, v := range strings.ToLower(subject) {
		if ok := strings.Contains(candidate, string(v)); !ok {
			return false
		}
		candidate = strings.Replace(candidate, string(v), "", 1)
	}

	return true
}
