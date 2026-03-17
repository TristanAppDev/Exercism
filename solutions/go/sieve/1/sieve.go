package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	bools := make([]bool, limit+1)
	primes := make([]int, 0, limit-1)

	for prime := 2; prime <= limit; {
		for i := prime + prime; i <= limit; i += prime {
			bools[i] = true
		}
		for prime++; prime <= limit && bools[prime]; prime++ {
		}
	}

	for i := 2; i <= limit; i++ {
		if !bools[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
