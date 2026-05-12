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

type AdminPatch struct {
	Email     Nullable[string]
	FullName  Nullable[string]
	AdminType Nullable[Role]
}

func NewAdminPatch(email, fullName Nullable[string], adminType Nullable[Role]) AdminPatch {
	return AdminPatch{
		Email:     email,
		FullName:  fullName,
		AdminType: adminType,
	}
}

func (p *AdminPatch) Validate() error {
	if p.Email.Set && p.Email.Value == nil {
		return fmt.Errorf("email can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("fullname can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.AdminType.Set && p.AdminType.Value == nil {
		return fmt.Errorf("admin type can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Email.Set && p.Email.Value != nil {
		if err := core_validation.ValidateEmail(*p.Email.Value); err != nil {
			return fmt.Errorf("invalid email in patch: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if p.FullName.Set && p.FullName.Value != nil {
		fullNameLength := len([]rune(*p.FullName.Value))
		if err := core_validation.ValidateIntInBounds(fullNameLength, MinFullNameLength, MaxFullNameLength); err != nil {
			return fmt.Errorf("invalid fullname length in patch: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if p.AdminType.Set && p.AdminType.Value != nil {
		if !IsRoleAdmin(*p.AdminType.Value) {
			return fmt.Errorf("unknown admin type in patch: %s: %w", *p.AdminType.Value, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

// ApplyPatch применяет патч к администратору
func (a *Admin) ApplyPatch(patch AdminPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid admin patch: %w", err)
	}

	tmp := *a

	if patch.Email.Set {
		tmp.Email = *patch.Email.Value
	}

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.AdminType.Set {
		tmp.AdminType = *patch.AdminType.Value
	}

	if err := ValidateAdmin(tmp.Email, tmp.FullName, "dummydummydummy", tmp.AdminType); err != nil {
		return fmt.Errorf("invalid patched admin: %w", err)
	}

	*a = tmp

	return nil
}
