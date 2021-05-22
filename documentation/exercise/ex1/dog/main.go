package dog

import (
	"errors"
)

const humanToDogMultiplier = 7

// Years takes in the current age of the dogs age and
// returns the human age and error
func Years(dogYears int) (int, error)  {
	if dogYears < 1 {
		return 0, errors.New("dog years need to be greater than 0")
	}
	return dogYears * humanToDogMultiplier, nil
}