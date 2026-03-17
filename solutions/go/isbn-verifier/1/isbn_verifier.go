package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN returns true if isbn is a valid ISBN-10
func IsValidISBN(isbn string) bool {
	// if isbn empty return false
	if len(isbn) < 10 {
		return false
	}

	// clean string and convert to slice of string
	cleanIsbn := strings.ReplaceAll(isbn, "-", "")
	chars := strings.Split(cleanIsbn, "")
	length := len(chars)

	// check len
	if length != 10 {
		return false
	}

	// check if all chars are numbers
	for _, v := range chars[:length-1] {
		_, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
	}

	// change last digit to 10 if x
	if chars[length-1] == "X" {
		chars[length-1] = "10"
	}

	// calulate sum
	sum := 0
	for k, v := range chars {
		num, _ := strconv.Atoi(v)
		sum += num * (10 - k)
	}

	// return true if sum is divisible by 11
	return sum%11 == 0
}
