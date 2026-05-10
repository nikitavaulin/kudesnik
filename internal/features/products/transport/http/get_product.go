package products_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
)

type GetProductResponseDTO ProductDTOResponse

func (h *ProductsHTTPHandler) GetProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke get product handler")

	category := domain.GetCategoryName(core_http_request.GetStringQueryParam(r, "category"))

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	product, err := h.productsService.GetProduct(ctx, productID, category)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get product by ID")
		return
	}

	responseHandler.JSONResponse(product, http.StatusOK)
}
