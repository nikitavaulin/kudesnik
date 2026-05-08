package products_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type UpdateProductVisabilityRequestDTO struct {
	IsVisible bool `json:"is_visible"`
}

func (h *ProductsHTTPHandler) UpdateProductVisability(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch product category handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	var dto UpdateProductVisabilityRequestDTO

	if err := core_http_request.DecodeRequest(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	if err := h.productsService.UpdateProductVisability(ctx, productID, dto.IsVisible); err != nil {
		responseHandler.ErrorResponse(err, "failed to patch product category")
		return
	}

	responseHandler.NoContentResponse()
}

type UpdateProductsVisabilityRequestDTO struct {
	IsVisible   bool `json:"is_visible"`
	ProductsIDs []ProductIdDTORequest
}

func (h *ProductsHTTPHandler) UpdateProductsVisability(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke update product visibility handler")

	var dto UpdateProductsVisabilityRequestDTO

	if err := core_http_request.DecodeRequest(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	productsIDs := UUIDsFromDTOs(dto.ProductsIDs...)

	if err := h.productsService.UpdateProductsVisability(ctx, productsIDs, dto.IsVisible); err != nil {
		responseHandler.ErrorResponse(err, "failed to update product visibility")
		return
	}

	responseHandler.NoContentResponse()
}
