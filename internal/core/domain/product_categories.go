package domain

import "github.com/google/uuid"

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
