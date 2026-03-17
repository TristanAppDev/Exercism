package romannumerals

import (
	"fmt"
	"strings"
)

type rule struct {
	arabic int
	roman  string
}

var rules = []rule{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {

	if input < 1 || input > 3000 {
		return "", fmt.Errorf("input must be between 1 and 3000")
	}

	var result strings.Builder

	for _, rule := range rules {
		for input >= rule.arabic {
			result.WriteString(rule.roman)
			input -= rule.arabic
		}
	}

	return result.String(), nil
}
