package lsproduct

import "fmt"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, fmt.Errorf("span is negative: %d", span)
	}
	if len(digits) < span {
		return 0, fmt.Errorf("span > len(%s): %d < %d", digits, len(digits), span)
	}

	numbers := make([]int64, len(digits))
	for i, r := range digits {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("%q contains non-digits", digits)
		}
		numbers[i] = int64(r - '0')
	}

	res := int64(0)
	for i, last := 0, len(numbers)-span+1; i < last; i++ {
		p := int64(1)
		for _, d := range numbers[i : i+span] {
			p *= d
		}
		if p > res {
			res = p
		}
	}
	return res, nil
}
