package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError type here.
type SillyNephewError struct {
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	}

	fodderAmount, err := weightFodder.FodderAmount()

	switch {
	case err != nil && err != ErrScaleMalfunction:
		return 0, err
	case fodderAmount < 0:
		return 0, errors.New("negative fodder")
	case err != nil && err == ErrScaleMalfunction && fodderAmount > 0:
		if fodderAmount > float64(cows) {
			return fodderAmount, nil
		} else {
			return 1, nil
		}
	default:
		return fodderAmount / float64(cows), nil
	}
}
