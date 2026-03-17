package parsinglogfiles

import (
	"regexp"
	"strings"
)

var prefixes = []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

func IsValidLine(text string) bool {
	return Any(prefixes, func(prefix string) bool {
		return strings.HasPrefix(text, prefix)
	})
}

func SplitLogLine(text string) []string {
	str := `<[*~=-]*>`
	reg := regexp.MustCompile(str)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	str := `(?i)".*password.*"`
	reg := regexp.MustCompile(str)
	for _, line := range lines {
		if reg.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line\d+`)
	return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`\sUser\s+(\S+)`)
	for i, line := range lines {
		match := reg.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		lines[i] = line
	}
	return lines
}

// Any function for practise of collection functions
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
