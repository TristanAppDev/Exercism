// Package bob should have a package comment that summarizes what it's about.
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
	case IsYellingQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case IsQuestion(remark):
		return "Sure."
	case IsYelling(remark):
		return "Whoa, chill out!"
	case IsSilence(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

// IsQuestion returns true if the remark ends with a question mark.
func IsQuestion(remark string) bool {
	remark = strings.TrimSpace(remark)
	return strings.HasSuffix(remark, "?")
}

// IsYelling returns true if the remark is in all caps.
func IsYelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}

// IsYellingQuestion returns true if the remark is in all caps and ends with a question mark.
func IsYellingQuestion(remark string) bool {
	return IsYelling(remark) && IsQuestion(remark)
}

// IsSilence returns true if the remark is empty.
func IsSilence(remark string) bool {
	return strings.TrimSpace(remark) == ""
}
