package domain

import (
	"fmt"

	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

type Customer struct {
	PhoneNumber string
	FullName    *string
}

func NewCustomer(phone string, fullname *string) (*Customer, error) {
	if err := validateCustomer(phone, fullname); err != nil {
		return nil, fmt.Errorf("invalid customer: %w", err)
	}
	return &Customer{
		PhoneNumber: phone,
		FullName:    fullname,
	}, nil
}

func validateCustomer(phone string, fullname *string) error {
	if err := core_validation.ValidatePhoneNumber(phone); err != nil {
		return err
	}
	if fullname != nil {
		return ValidateFullName(*fullname)
	}
	return nil
}
