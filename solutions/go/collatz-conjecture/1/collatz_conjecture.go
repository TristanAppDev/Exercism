package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("0 and negativ numbers not allowed")
	}

	for counter := 0; ; counter++ {
		if n == 1 {
			return counter, nil
		}

		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
}
