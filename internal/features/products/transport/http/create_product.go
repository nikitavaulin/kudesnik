package products_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type CreateProductResponse ProductDTOResponse

func (h *ProductsHTTPHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(logger, rw)

	logger.Debug("invoke create product handler")

	category := core_http_request.GetStringQueryParam(r, "category")
	product := domain.GetProductEmptyInstance(category)

	if err := core_http_request.DecodeRequest(r, &product); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	product, err := h.productsService.CreateProduct(ctx, product)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create product")
		return
	}

	responseHandler.JSONResponse(product, http.StatusCreated)
}
