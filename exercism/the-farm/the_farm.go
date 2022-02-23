package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	NoOfCows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.NoOfCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, ok := weightFodder.FodderAmount()
	nephew := &SillyNephewError{NoOfCows: cows}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if ok == nil {
		if cows < 0 {
			return 0, errors.New(nephew.Error())
		}

		if amount < 0 {
			return 0, errors.New("negative fodder")
		}
	}

	if ok != nil {

		if ok.Error() == ErrScaleMalfunction.Error() {
			if amount > 0 {
				return amount * 2 / float64(cows), nil
			}
			return 0, errors.New("negative fodder")
		}

		if nephew.NoOfCows < 0 {

			return 0, errors.New(nephew.Error())
		}

		return 0, errors.New("non-scale error")

	}

	return amount / float64(cows), nil
}
