package products_transport_http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type CreateProductRequest struct {
	ProductName string     `json:"product_name" validate:"required,min=3,max=100"`
	Price       float64    `json:"price"`
	Description *string    `json:"description"`
	CategoryID  uuid.UUID  `json:"category_id" validate:"required"`
	ProducerID  *uuid.UUID `json:"producerId"`
}

type CreateProductResponse ProductDTOResponse

func (h *ProductsHTTPHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(logger, rw)

	logger.Debug("invoke create product handler")

	var productRequestDTO CreateProductRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &productRequestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	product := domainFromDTO(productRequestDTO)

	product, err := h.productsService.CreateProduct(ctx, product)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create product")
		return
	}

	productResponseDTO := CreateProductResponse(productDtoFromDomain(product))
	responseHandler.JSONResponse(productResponseDTO, http.StatusCreated)
}

func domainFromDTO(dto CreateProductRequest) domain.BaseProduct {
	return *domain.NewProductUninitialized(
		dto.ProductName,
		dto.Price,
		dto.Description,
		dto.CategoryID,
		dto.ProducerID,
	)
}
