package products_transport_http

import (
	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductIdDTORequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
}

func UUIDsFromDTOs(dtos ...ProductIdDTORequest) []uuid.UUID {
	IDs := make([]uuid.UUID, len(dtos))
	for i, dto := range dtos {
		IDs[i] = dto.ProductID
	}
	return IDs
}

type ProductDTOResponse struct {
	ID          string  `json:"id"`
	Version     int     `json:"version"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	IsVisible   bool    `json:"is_visible"`
	CategoryID  string  `json:"category_id"`
	ProducerID  *string `json:"producer_id"`
}

func productDtoFromDomain(product domain.BaseProduct) ProductDTOResponse {
	var producerID *string
	if product.ProducerID != nil {
		id := *product.ProducerID
		idStr := id.String()
		producerID = &idStr
	}

	return ProductDTOResponse{
		ID:          product.ID.String(),
		Version:     product.Version,
		ProductName: product.ProductName,
		Price:       product.Price,
		Description: product.Description,
		IsVisible:   product.IsVisible,
		CategoryID:  product.CategoryID.String(),
		ProducerID:  producerID,
	}
}

func productsDtoFromDomain(products ...domain.BaseProduct) []ProductDTOResponse {
	dtos := make([]ProductDTOResponse, len(products))
	for i, product := range products {
		dto := productDtoFromDomain(product)
		dtos[i] = dto
	}
	return dtos
}
