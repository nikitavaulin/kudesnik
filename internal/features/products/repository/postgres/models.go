package products_repository_postges

import (
	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductModel struct {
	ID          uuid.UUID
	Version     int
	ProductName string
	Price       float64
	Description *string
	IsVisible   bool
	CategoryID  uuid.UUID
	ProducerID  *uuid.UUID
}

func productDomainFromModel(m ProductModel) domain.BaseProduct {
	return *domain.NewProduct(
		m.ID, m.Version,
		m.ProductName, m.Price, m.Description,
		m.IsVisible, m.CategoryID, m.ProducerID,
	)
}

func productsDomainFromModels(models ...ProductModel) []domain.BaseProduct {
	products := make([]domain.BaseProduct, len(models))
	for i, model := range models {
		product := productDomainFromModel(model)
		products[i] = product
	}
	return products
}
