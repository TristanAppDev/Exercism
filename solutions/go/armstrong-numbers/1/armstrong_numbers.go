package armstrong

func IsNumber(n int) bool {
	if n < 10 {
		return true
	}

	if n < 100 {
		return false
	}

	exp := length(n)
	remnant := n
	sum := 0

	for remnant > 0 {
		sum += pow(remnant%10, exp)
		remnant /= 10
	}

	return sum == n
}

// pow returns x^y faster than math.Pow(x, y) because it avoids casting to float64 and back to int.
func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	result := x
	for i := 2; i <= y; i++ {
		result *= x
	}
	return result
}

// length returns the count of digits in the int.
func length(n int) int {
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}
