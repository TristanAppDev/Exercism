package isbn

// IsValidISBN returns true if isbn is a valid ISBN-10
func IsValidISBN(isbn string) bool {
	sum := 0
	counter := 10 // should be 0 if cleaned isbn has correct length
	for i, r := range isbn {
		switch {
		case r >= '0' && r <= '9':
			sum += Rtoi(r) * counter
		case r == 'X' && i == len(isbn)-1:
			sum += 10 * counter
		case r == '-':
			continue
		default:
			return false
		}
		counter--
	}
	return counter == 0 && sum%11 == 0
}

func Rtoi(r rune) int {
	return int(r - '0')
}
