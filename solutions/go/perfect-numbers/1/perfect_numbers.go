package perfect

import "errors"

// Define the Classification type here.
type Classification string

const (
	ClassificationDeficient Classification = "deficient"
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
)

var (
	ErrOnlyPositive = errors.New("Only positive numbers are allowed")
)

func Classify(n int64) (Classification, error) {

	if n <= 0 {
		return "", ErrOnlyPositive
	}

	sum := AliquotSum(n)

	if sum < n {
		return ClassificationDeficient, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	} else {
		return ClassificationPerfect, nil
	}
}

func AliquotSum(n int64) int64 {
	var sum int64
	for i := int64(1); i < n/2+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
