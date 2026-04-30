package domain

import (
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

const (
	MinProductCategoryNameLength = 3
	MaxProductCategoryNameLength = 60
)

type ProductCategory struct {
	ID                uuid.UUID
	CategoryName      string
	InstallationPrice float64
}

func NewProductCategoryUninitialized(name string, installationPrice float64) *ProductCategory {
	return NewProductCategory(UnizitializedID, name, installationPrice)
}

func NewProductCategory(id uuid.UUID, name string, installationPrice float64) *ProductCategory {
	return &ProductCategory{
		ID:                id,
		CategoryName:      name,
		InstallationPrice: installationPrice,
	}
}

func (c *ProductCategory) Validate() error {
	nameLength := len([]rune(c.CategoryName))
	err := core_validation.ValidateIntInBounds(nameLength, MinProductCategoryNameLength, MaxProductCategoryNameLength)
	if err != nil {
		return fmt.Errorf(
			"invalid category name: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	if c.InstallationPrice < 0 {
		return fmt.Errorf(
			"installation price should be greater than 0, got: %v: %w",
			c.InstallationPrice,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type ProductCategoryPatch struct {
	CategoryName      Nullable[string]
	InstallationPrice Nullable[float64]
}

func (c *ProductCategoryPatch) Validate() error {
	if c.CategoryName.Set && c.CategoryName.Value == nil {
		return fmt.Errorf("CategoryName can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (c *ProductCategory) ApplyPatch(patch ProductCategoryPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid product category patch: %w", err)
	}

	tmp := *c

	if patch.CategoryName.Set {
		tmp.CategoryName = *patch.CategoryName.Value
	}

	if patch.InstallationPrice.Set {
		tmp.InstallationPrice = *patch.InstallationPrice.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("ivalid patched product category: %w", err)
	}

	*c = tmp

	return nil
}
