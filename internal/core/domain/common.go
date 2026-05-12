package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

const (
	MinFullNameLength = 1
	MaxFullNameLength = 60
)

func ValidateFullName(fullname string) error {
	if fullname == "" {
		return fmt.Errorf("fullname cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidateIntInBounds(len(fullname), MinFullNameLength, MaxFullNameLength); err != nil {
		return fmt.Errorf("invalid fullname length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}
