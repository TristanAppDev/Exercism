package phonenumber

import (
	"errors"
)

var invalidPhoneNumberErr = errors.New("phone number is invalid")

func Number(phoneNumber string) (string, error) {
	digits, err := extractDigits(phoneNumber)
	if err != nil {
		return "", err
	}

	numDigits := len(digits)

	if number, ok := isValidPhoneNumber(digits, numDigits); ok {
		return string(number), nil
	}

	return "", invalidPhoneNumberErr
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + number[:3] + ") " + number[3:6] + "-" + number[6:], nil
}

func isValidPhoneNumber(digits []rune, numDigits int) ([]rune, bool) {
	if numDigits == 10 && digits[0] >= '2' && digits[3] >= '2' {
		return digits, true
	}
	if numDigits == 11 && digits[0] == '1' && digits[1] >= '2' && digits[4] >= '2' {
		return digits[1:], true
	}
	return []rune{}, false
}

func extractDigits(phoneNumber string) ([]rune, error) {
	digits := make([]rune, 0, len(phoneNumber))
	for _, r := range phoneNumber {
		switch r {
		case ' ', '+', '(', ')', '-', '.':
		default:
			if r < '0' || r > '9' {
				return []rune{}, invalidPhoneNumberErr
			}
			digits = append(digits, r)
		}
	}
	return digits, nil
}
