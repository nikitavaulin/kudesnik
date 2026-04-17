package product_categories_repository

import "github.com/google/uuid"

type ProductCategoriesModel struct {
	ID                uuid.UUID
	CategoryName      string
	InstallationPrice float64
}
