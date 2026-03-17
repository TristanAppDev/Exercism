package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("numbers below 1 are not allowed")
	}

	for c := 0; ; c++ {
		if n == 1 {
			return c, nil
		}

		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
}
