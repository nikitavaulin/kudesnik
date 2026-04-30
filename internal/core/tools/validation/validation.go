package core_validation

import "fmt"

func ValidateIntInBounds(number, minValue, maxValue int) error {
	if !(minValue <= number && number <= maxValue) {
		return fmt.Errorf(
			"should be between %d and %d, got: %d",
			minValue,
			maxValue,
			number,
		)
	}
	return nil
}
