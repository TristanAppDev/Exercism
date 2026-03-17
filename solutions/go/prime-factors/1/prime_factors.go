package prime

import (
	"math"
)

func Factors(n int64) []int64 {

	if n == 1 {
		return []int64{}
	}

	if n == 2 {
		return []int64{2}
	}

	result := make([]int64, 0)

	for i := 2; int64(i) <= n; i++ {

		if !isPrime(int64(i)) {
			continue
		}

		for n%int64(i) == 0 {
			result = append(result, int64(i))
			n = n / int64(i)
		}
	}

	return result
}

func isPrime(number int64) bool {

	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(number)))+1; i++ {
		if number%int64(i) == 0 {
			return false
		}
	}
	return true
}
