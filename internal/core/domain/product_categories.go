package domain

import (
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
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
	if !validateIntInBounds(nameLength, MinProductCategoryNameLength, MaxProductCategoryNameLength) {
		return fmt.Errorf(
			"product category should be between %d and %d, got: %d: %w",
			MinProductCategoryNameLength, MaxProductCategoryNameLength,
			nameLength,
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

func validateIntInBounds(number, minValue, maxValue int) bool {
	return minValue <= number && number <= maxValue
}
