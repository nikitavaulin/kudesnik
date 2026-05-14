package products_repository_postges

import (
	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductModel struct {
	ID           uuid.UUID
	Version      int
	ProductName  string
	Price        float64
	Description  *string
	IsVisible    bool
	CategoryCode string
	ProducerID   *uuid.UUID
}

func productDomainFromModel(m ProductModel) domain.ProductBase {
	return *domain.NewProduct(
		m.ID, m.Version,
		m.ProductName, m.Price, m.Description,
		m.IsVisible, m.CategoryCode, m.ProducerID,
		nil, nil,
	)
}

func productsDomainFromModels(models ...ProductModel) []domain.ProductBase {
	products := make([]domain.ProductBase, len(models))
	for i, model := range models {
		product := productDomainFromModel(model)
		products[i] = product
	}
	return products
}
