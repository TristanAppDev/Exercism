package prime

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	count := 0
	for i := 1; ; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i, true
			}
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
