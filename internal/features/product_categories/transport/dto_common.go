package product_categories_transport_http

import "github.com/nikitavaulin/kudesnik/internal/core/domain"

type ProductCategoryDTOResponse struct {
	CategoryCode      string  `json:"category_code"`
	CategoryName      string  `json:"category_name"`
	InstallationPrice float64 `json:"installation_price"`
}

func dtoFromDomain(category domain.ProductCategory) ProductCategoryDTOResponse {
	return ProductCategoryDTOResponse{
		CategoryCode:      category.CategoryCode,
		CategoryName:      category.CategoryName,
		InstallationPrice: category.InstallationPrice,
	}
}

func dtosFromDomains(categories []domain.ProductCategory) []ProductCategoryDTOResponse {
	categoriesDTO := make([]ProductCategoryDTOResponse, len(categories))
	for i, category := range categories {
		categoriesDTO[i] = dtoFromDomain(category)
	}
	return categoriesDTO
}
