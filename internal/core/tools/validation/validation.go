package core_validation

import (
	"fmt"
	"regexp"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
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

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("email dismatch pattern")
	}

	return nil
}

const (
	MinPhoneLength = 10
	MaxPhoneLength = 15
)

func ValidatePhoneNumber(phone string) error {
	if phone == "" {
		return fmt.Errorf("phone number cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := ValidateIntInBounds(len(phone), MinPhoneLength, MaxPhoneLength); err != nil {
		return fmt.Errorf("invalid phone length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	phoneRegex := regexp.MustCompile(`^\+[0-9]+$`)
	if !phoneRegex.MatchString(phone) {
		return fmt.Errorf("dismatch phone number pattern: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
