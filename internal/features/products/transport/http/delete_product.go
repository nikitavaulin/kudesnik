package products_transport_http

import (
	"net/http"

	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

func (h *ProductsHTTPHandler) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete product handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	if err := h.productsService.DeleteProduct(ctx, productID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete product")
		return
	}

	responseHandler.NoContentResponse()
}

func (h *ProductsHTTPHandler) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke delete products handler")

	var productIDsDTO []ProductIdDTORequest

	if err := core_http_request.DecodeRequest(r, &productIDsDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	productIDs := UUIDsFromDTOs(productIDsDTO...)

	if err := h.productsService.DeleteProducts(ctx, productIDs); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete products")
		return
	}

	responseHandler.NoContentResponse()
}
