package domain

import (
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

type Admin struct {
	ID           uuid.UUID
	Email        string
	FullName     string
	PasswordHash string
	AdminType    Role
}

func NewAdmin(email, fullName, passwordHash string, adminType Role) *Admin {
	return &Admin{
		ID:           uuid.New(),
		Email:        email,
		FullName:     fullName,
		PasswordHash: passwordHash,
		AdminType:    adminType,
	}
}

const (
	MinFullNameLength = 1
	MaxFullNameLength = 60
	MinPasswordLength = 8
	MaxPasswordLength = 72
)

func IsRoleAdmin(role Role) bool {
	validRoles := map[Role]bool{
		AdminRole:          true,
		ManagerRole:        true,
		DismissedAdminRole: true,
	}
	return validRoles[role]
}

func ValidateAdmin(email, fullName, password string, adminType Role) error {
	if err := core_validation.ValidateEmail(email); err != nil {
		return fmt.Errorf("invalid email: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if fullName == "" {
		return fmt.Errorf("fullname cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidateIntInBounds(len(fullName), MinFullNameLength, MaxFullNameLength); err != nil {
		return fmt.Errorf("invalid fullname length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if password == "" {
		return fmt.Errorf("password cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidateIntInBounds(len(password), MinPasswordLength, MaxPasswordLength); err != nil {
		return fmt.Errorf("invalid password length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if adminType == "" {
		return fmt.Errorf("admin type cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if !IsRoleAdmin(adminType) {
		return fmt.Errorf("unknown admin type: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
