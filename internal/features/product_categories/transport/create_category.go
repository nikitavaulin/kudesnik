package product_categories_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type CreateCategoryRequest struct {
	CategoryName      string  `json:"category_name" validate:"required,min=3,max=60"`
	InstallationPrice float64 `json:"installation_price"`
}

type CreateCategoryResponse struct {
	ID                string  `json:"id"`
	CategoryName      string  `json:"category_name"`
	InstallationPrice float64 `json:"installation_price"`
}

func (h *ProductCategoryHTTPHandler) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke create product category handler")

	var requestDTO CreateCategoryRequest
	if err := core_http_request.DecodeAndValidateRequest(r, requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	category := domainFromDTO(requestDTO)

	category, err := h.categoryService.CreateProductCategory(ctx, category)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create product category")
		return
	}

	responseDTO := dtoFromDomain(category)

	responseHandler.JSONResponse(responseDTO, http.StatusCreated)
}

func domainFromDTO(dto CreateCategoryRequest) domain.ProductCategory {
	return *domain.NewProductCategoryUninitialized(dto.CategoryName, dto.InstallationPrice)
}

func dtoFromDomain(category domain.ProductCategory) CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:                category.ID.String(),
		CategoryName:      category.CategoryName,
		InstallationPrice: category.InstallationPrice,
	}
}
