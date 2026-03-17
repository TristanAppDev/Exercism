package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError type here.
type SillyNephewError struct {
	Cows int
}

func (sillyNephew *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sillyNephew.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	fodderAmount, err := weightFodder.FodderAmount()

	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if fodderAmount < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, &SillyNephewError{Cows: cows}
	}
	if err == ErrScaleMalfunction {
		fodderAmount *= 2
	}

	return fodderAmount / float64(cows), nil
}
