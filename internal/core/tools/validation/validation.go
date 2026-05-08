package core_validation

import (
	"fmt"
)

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

func ValidateLimitOffset(limit, offset *int) error {
	if limit != nil && *limit < 0 {
		return fmt.Errorf("limit must be non-negative")
	}

	if offset != nil && *offset < 0 {
		return fmt.Errorf("offset must be non-negative")
	}

	return nil
}
