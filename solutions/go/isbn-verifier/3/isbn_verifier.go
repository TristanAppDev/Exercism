package isbn

// IsValidISBN returns true if isbn is a valid ISBN-10
func IsValidISBN(isbn string) bool {
	sum := 0
	counter := 10

	for i, r := range isbn {
		switch {
		case r >= '0' && r <= '9':
			sum += int(r-'0') * counter
			counter--
		case r == 'X' && i == len(isbn)-1:
			sum += 10 * counter
			counter--
		case r == '-':
			continue
		default:
			return false
		}
	}
	return sum%11 == 0 && counter == 0
}
