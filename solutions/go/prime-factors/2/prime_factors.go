package prime

func Factors(n int64) []int64 {
	result := make([]int64, 0)
	var factor int64 = 2

	for factor*factor <= n {
		for n%factor == 0 {
			result = append(result, factor)
			n = n / factor
		}
		factor++
	}

	if n > 1 {
		result = append(result, n)
	}

	return result
}
